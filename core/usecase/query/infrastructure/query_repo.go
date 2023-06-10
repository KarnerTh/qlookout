package infrastructure

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/core/usecase/query"
)

type queryRepo struct {
	db *sql.DB
}

func NewQueryRepo(db *sql.DB) query.QueryRepo {
	return queryRepo{db: db}
}

func (q queryRepo) Query(queryString string) (query.QueryResult, error) {
	rows, err := q.db.Query(queryString)
	if err != nil {
		return query.QueryResult{}, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return query.QueryResult{}, err
	}

	var rowDataList []map[string]any

	for rows.Next() {
		columnValues := make([]any, len(columns))
		columnPointers := make([]any, len(columns))
		for i := range columns {
			columnPointers[i] = &columnValues[i]
		}

		if err = rows.Scan(columnPointers...); err != nil {
			return query.QueryResult{}, err
		}

		rowValue := make(map[string]any)
		for i, column := range columns {
			rowValue[column] = *columnPointers[i].(*any)
		}

		rowDataList = append(rowDataList, rowValue)
	}

	return query.QueryResult{
		Columns: columns,
		Rows:    rowDataList,
	}, nil
}
