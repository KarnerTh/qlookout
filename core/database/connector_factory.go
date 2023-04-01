package database

import (
	"database/sql"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

type connectorFactory struct{}

type ConnectorFactory interface {
	NewConnector(dataSource string) (*sql.DB, error)
}

func NewConnectorFactory() ConnectorFactory {
	return connectorFactory{}
}

func (connectorFactory) NewConnector(dataSource string) (*sql.DB, error) {
	if dataSource == "" {
		return nil, errors.New("No data source provided")
	}

	switch {
	case strings.HasPrefix(dataSource, "sqlite3"):
		return open(openSqlite, Sqlite3, dataSource)
	case strings.HasPrefix(dataSource, "postgresql") || strings.HasPrefix(dataSource, "postgres"):
		return open(openPostgres, Postgres, dataSource)
	case strings.HasPrefix(dataSource, "mysql"):
		return open(openMysql, MySql, dataSource)
	case strings.HasPrefix(dataSource, "sqlserver"):
		return open(openMssql, MsSql, dataSource)
	default:
		return nil, errors.New("could not create connector for db")
	}
}

func open(openFunc func(driverName DbType, dataSource string) (*sql.DB, error), drivername DbType, dataSource string) (*sql.DB, error) {
	db, err := openFunc(drivername, dataSource)
	if err != nil {
		return nil, err
	}

	err = verifyConnection(db, drivername)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func verifyConnection(db *sql.DB, driverName DbType) error {
	err := db.Ping()
	if err != nil {
		log.WithError(err).Errorf("Could not verify db connection %s", driverName)
		return err
	}

	log.Debugf("Connected to %s db successfully", driverName)
	return nil
}
