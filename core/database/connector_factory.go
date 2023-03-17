package database

import (
	"database/sql"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

type connectorFactory struct{}

type ConnectorFactory interface {
	NewConnector(connectionString string) (*sql.DB, error)
}

func NewConnectorFactory() ConnectorFactory {
	return connectorFactory{}
}

func (connectorFactory) NewConnector(connectionString string) (*sql.DB, error) {
	switch {
	case strings.HasPrefix(connectionString, "sqlite3"):
		return open(openSqlite, Sqlite3, connectionString)
	case strings.HasPrefix(connectionString, "postgresql") || strings.HasPrefix(connectionString, "postgres"):
		return open(openPostgres, Postgres, connectionString)
	case strings.HasPrefix(connectionString, "mysql"):
		return open(openMysql, MySql, connectionString)
	case strings.HasPrefix(connectionString, "sqlserver"):
		return open(openMssql, MsSql, connectionString)
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
