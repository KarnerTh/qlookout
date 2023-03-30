package review

type ColumnType string

const (
	Text  ColumnType = "text"
	Int   ColumnType = "int"
	Float ColumnType = "float"
)

type ReviewRule struct {
	Id           int
	LookoutId    int
	ColumnName   string
	ColumnType   ColumnType
	RowIndex     int
	ExactValue   string
	GreaterThan  string
	LessThan     string
	ShouldBeNull bool
}
