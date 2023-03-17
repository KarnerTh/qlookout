package query

type QueryRepo interface {
	Query(queryString string) (QueryResult, error)
}
