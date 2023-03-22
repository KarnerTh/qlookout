package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/KarnerTh/query-lookout/usecase/lookout"
)

type lookoutRepo struct {
	db *sql.DB
}

func NewLookoutRepo(db *sql.DB) lookout.LookoutRepo {
	return &lookoutRepo{db: db}
}

func (r lookoutRepo) GetConfigs() ([]lookout.LookoutConfig, error) {
	rows, err := r.db.Query("select id, name, query, cron from lookout")
	if err != nil {
		return nil, err
	}

	var lookouts []lookout.LookoutConfig
	for rows.Next() {
		var l lookout.LookoutConfig
		if err = rows.Scan(&l.Id, &l.Name, &l.Query, &l.Cron); err != nil {
			return nil, err
		}

		lookouts = append(lookouts, l)
	}

	return lookouts, nil
}

func (r lookoutRepo) GetConfig(id int) (*lookout.LookoutConfig, error) {
	rows, err := r.db.Query("select id, name, query, cron from lookout where id = ?", id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("No lookout config found for id %d", id)
	}

	var config lookout.LookoutConfig
	if err = rows.Scan(&config.Id, &config.Name, &config.Query, &config.Cron); err != nil {
		return nil, err
	}

	// only one config should be present with the id
	if rows.Next() {
		return nil, fmt.Errorf("Multiple config items found with id %d", id)
	}

	return &config, nil
}
