package graphql

type reviewRuleCreateModel struct {
	LookoutId    int32
	ColumnName   string
	ColumnType   string
	RowIndex     int32
	ExactValue   *string
	GreaterThan  *string
	LessThan     *string
	ShouldBeNull *bool
}
