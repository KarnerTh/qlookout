package database

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func openSqlite(driverName DbType, dataSource string) (*sql.DB, error) {
	connectionString := strings.ReplaceAll(dataSource, "sqlite3://", "file:")
	db, err := sql.Open(driverName.String(), connectionString)
	if err != nil {
		log.WithError(err).Errorf("Could not open database %s with connection string %s", driverName, dataSource)
		return nil, err
	}

	return db, nil
}
