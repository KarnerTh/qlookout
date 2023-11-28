package infrastructure

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KarnerTh/qlookout/core/usecase/lookout"
)

type lookoutRepo struct {
	db *sql.DB
}

func NewLookoutRepo(db *sql.DB) lookout.LookoutRepo {
	return lookoutRepo{db: db}
}

func (r lookoutRepo) Get() ([]lookout.LookoutConfig, error) {
	rows, err := r.db.Query("select id, name, query, cron, notify_local, notify_mail from lookout")
	if err != nil {
		return nil, err
	}

	var lookouts []lookout.LookoutConfig
	for rows.Next() {
		var config lookout.LookoutConfig
		if err = rows.Scan(&config.Id, &config.Name, &config.Query, &config.Cron, &config.NotifyLocal, &config.NotifyMail); err != nil {
			return nil, err
		}

		lookouts = append(lookouts, config)
	}

	return lookouts, nil
}

func (r lookoutRepo) GetById(id int) (*lookout.LookoutConfig, error) {
	rows, err := r.db.Query("select id, name, query, cron, notify_local, notify_mail from lookout where id = ?", id)

	// if rows.Next is not called til it returns false, the rows are not automatically closed
	// source: https://pkg.go.dev/database/sql#Rows.Close
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("No lookout config found for id %d", id)
	}

	var config lookout.LookoutConfig
	if err = rows.Scan(&config.Id, &config.Name, &config.Query, &config.Cron, &config.NotifyLocal, &config.NotifyMail); err != nil {
		return nil, err
	}

	// only one config should be present with the id
	if rows.Next() {
		return nil, fmt.Errorf("Multiple config items found with id %d", id)
	}

	return &config, nil
}

func (r lookoutRepo) Create(data lookout.LookoutConfigCreate) (*lookout.LookoutConfig, error) {
	var id int
	err := r.db.QueryRow(`
insert into lookout(name, query, cron, notify_local, notify_mail)
values (?, ?, ?, ?, ?)
returning id
    `, data.Name, data.Query, data.Cron, data.NotifyLocal, data.NotifyMail).Scan(&id)

	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r lookoutRepo) Update(id int, data lookout.LookoutConfigUpdate) (*lookout.LookoutConfig, error) {
	var updateProps []string
	var args []any

	if data.Name != nil {
		updateProps = append(updateProps, "name=?")
		args = append(args, data.Name)
	}
	if data.Cron != nil {
		updateProps = append(updateProps, "cron=?")
		args = append(args, data.Cron)
	}
	if data.Query != nil {
		updateProps = append(updateProps, "query=?")
		args = append(args, data.Query)
	}
	if data.NotifyLocal != nil {
		updateProps = append(updateProps, "notify_local=?")
		args = append(args, data.NotifyLocal)
	}
	if data.NotifyLocal != nil {
		updateProps = append(updateProps, "notify_mail=?")
		args = append(args, data.NotifyMail)
	}

	if len(updateProps) == 0 {
		// nothing to update
		return r.GetById(id)
	}

	query := "update lookout set " + strings.Join(updateProps, ",") + " where id = ?"
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r lookoutRepo) Delete(id int) (*lookout.LookoutConfig, error) {
	existing, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	query := "delete from lookout where id = ?"
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}

	return existing, nil
}
