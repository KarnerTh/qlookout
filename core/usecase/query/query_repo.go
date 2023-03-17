package query

type QueryRepo interface {
	Query(query string) (any, error)
}
