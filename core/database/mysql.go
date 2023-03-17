package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func openMysql(driverName DbType, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driverName.String(), dataSource)
	if err != nil {
		log.WithError(err).Errorf("Could not open database %s with connection string %s", driverName, dataSource)
		return nil, err
	}

	return db, nil
}
