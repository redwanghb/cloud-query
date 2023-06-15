package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
)

// CreateTableBatch migrates a table. It forms part of the writer.MixedBatchWriter interface.
func (c *Client) CreateTableBatch(ctx context.Context, messages []plugin.MessageCreateTable) error {
	tables, err := tablesFromMessages(messages)
	if err != nil {
		return err
	}
	pgTables, err := c.listPgTables(ctx, tables)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	tables = c.normalizeTables(tables, pgTables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(tables, pgTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range tables {
		tableName := table.Name
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}
		pgTable := pgTables.Get(tableName)
		if pgTable == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(ctx, table); err != nil {
				return err
			}
		} else {
			changes := table.GetChanges(pgTable)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, tableName); err != nil {
					return err
				}
				if err := c.createTableIfNotExist(ctx, table); err != nil {
					return err
				}
			}
		}
	}
	conn, err := c.conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()
	if err := conn.Conn().DeallocateAll(ctx); err != nil {
		return fmt.Errorf("failed to deallocate all prepared statements: %w", err)
	}
	return nil
}

func (c *Client) normalizeTable(table *schema.Table, pgTable *schema.Table) *schema.Table {
	normalizedTable := schema.Table{
		Name: table.Name,
	}
	for _, col := range table.Columns {
		if c.enabledPks() && col.PrimaryKey {
			col.NotNull = true
		} else {
			col.PrimaryKey = false
		}
		col.Type = c.PgToSchemaType(c.SchemaTypeToPg(col.Type))
		normalizedTable.Columns = append(normalizedTable.Columns, col)
	}

	if pgTable != nil && pgTable.PkConstraintName != "" {
		normalizedTable.PkConstraintName = pgTable.PkConstraintName
	}

	return &normalizedTable
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	tableName := table.Name
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if err := c.addColumn(ctx, tableName, change.Current); err != nil {
				return err
			}
		case schema.TableColumnChangeTypeRemove:
			continue
		default:
			panic("unknown change type")
		}
	}
	return nil
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			if change.Current.PrimaryKey || change.Current.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.PrimaryKey || change.Previous.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeUpdate:
			return false
		default:
			panic("unknown change type")
		}
	}
	return true
}

// normalize the requested schema to be compatible with what Postgres supports
func (c *Client) normalizeTables(tables schema.Tables, pgTables schema.Tables) schema.Tables {
	var result schema.Tables
	for _, table := range tables {
		pgTable := pgTables.Get(table.Name)
		if pgTable == nil {
			result = append(result, table)
		} else {
			result = append(result, c.normalizeTable(table, pgTable))
		}
	}
	return result
}

func (c *Client) nonAutoMigrableTables(tables schema.Tables, pgTables schema.Tables) ([]string, [][]schema.TableColumnChange) {
	var result []string
	var tableChanges [][]schema.TableColumnChange
	for _, t := range tables {
		pgTable := pgTables.Get(t.Name)
		if pgTable == nil {
			continue
		}
		changes := t.GetChanges(pgTable)
		if !c.canAutoMigrate(changes) {
			result = append(result, t.Name)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func (c *Client) dropTable(ctx context.Context, tableName string) error {
	c.logger.Info().Str("table", tableName).Msg("Dropping table")
	sql := "drop table " + tableName
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
	}
	return nil
}

func (c *Client) addColumn(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Column doesn't exist, creating")
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	sql := "alter table " + tableName + " add column " + columnName + " " + columnType
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	tName := table.Name
	tableName := pgx.Identifier{tName}.Sanitize()
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(tableName)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		pgType := c.SchemaTypeToPg(col.Type)
		columnName := pgx.Identifier{col.Name}.Sanitize()
		fieldDef := columnName + " " + pgType
		if col.Unique {
			fieldDef += " UNIQUE"
		}
		if col.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if c.enabledPks() && col.PrimaryKey {
			primaryKeys = append(primaryKeys, pgx.Identifier{col.Name}.Sanitize())
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(pgx.Identifier{tName + "_cqpk"}.Sanitize())
		sb.WriteString(" PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.conn.Exec(ctx, sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", tName, err)
	}
	return nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}