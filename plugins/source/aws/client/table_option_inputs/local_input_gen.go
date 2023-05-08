// Code generated by generator, DO NOT EDIT.
package table_option_inputs

type Granularity string
type DateInterval struct {
	End   *string
	Start *string
}
type MatchOption string
type CostCategoryValues struct {
	Key          *string
	MatchOptions *MatchOption
	Values       []string
}
type Dimension string
type DimensionValues struct {
	Key          *Dimension
	MatchOptions *MatchOption
	Values       []string
}
type TagValues struct {
	Key          *string
	MatchOptions *MatchOption
	Values       []string
}
type Expression struct {
	And            *Expression
	CostCategories *CostCategoryValues
	Dimensions     *DimensionValues
	Not            *Expression
	Or             *Expression
	Tags           *TagValues
}
type GroupDefinitionType string
type GroupDefinition struct {
	Key  *string
	Type *GroupDefinitionType
}
type GetCostAndUsageInput struct {
	Granularity   *Granularity
	Metrics       []string
	TimePeriod    *DateInterval
	Filter        *Expression
	GroupBy       *GroupDefinition
	NextPageToken *string
}