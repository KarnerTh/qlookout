package graphql

type reviewRuleUpdateModel struct {
	ColumnName   *string
	ColumnType   *string
	RowIndex     *int32
	ExactValue   *string
	GreaterThan  *string
	LessThan     *string
	ShouldBeNull *bool
}
