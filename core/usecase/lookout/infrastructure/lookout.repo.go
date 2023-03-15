package infrastructure

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/usecase/lookout"
)

type lookoutRepo struct {
	db *sql.DB
}

func NewLookoutRepo(db *sql.DB) lookout.LookoutRepo {
	return lookoutRepo{db: db}
}

func (r lookoutRepo) Get() ([]lookout.Lookout, error) {
	rows, err := r.db.Query("select id, name, query, cron from lookout")
	if err != nil {
		return nil, err
	}

	var lookouts []lookout.Lookout
	for rows.Next() {
		var l lookout.Lookout
		if err = rows.Scan(&l.Id, &l.Name, &l.Query, &l.Cron); err != nil {
			return nil, err
		}

		lookouts = append(lookouts, l)
	}

	return lookouts, nil
}
