package query

type QueryResult struct {
	Columns []string
	Rows    []map[string]any
}
