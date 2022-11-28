// Code generated by codegen; DO NOT EDIT.

package accounts

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AccountMembers() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_account_members",
		Resolver: fetchAccountMembers,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code"),
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "roles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Roles"),
			},
		},
	}
}
