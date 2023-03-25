package review

type ReviewRule struct {
	Id           int
	LookoutId    int
	ColumnName   string
	RowIndex     int
	ExactValue   string
	GreaterThan  string
	LessThan     string
	ShouldBeNull bool
}
