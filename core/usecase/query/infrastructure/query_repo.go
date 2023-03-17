package infrastructure

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/usecase/query"
)

type queryRepo struct {
	db *sql.DB
}

func NewQueryRepo(db *sql.DB) query.QueryRepo {
	return &queryRepo{db: db}
}

func (q queryRepo) Query(query string) (any, error) {
	rows, err := q.db.Query(query)
	if err != nil {
		return nil, err
	}

  // TODO: what about any?

	return rows, nil
}
