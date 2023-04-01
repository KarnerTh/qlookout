package database

import (
	"database/sql"
	"embed"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
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

//go:embed migrations/*.sql
var migrationsFS embed.FS

func MigrateInternalSqliteDb(db *sql.DB) error {
	// Source: https://github.com/golang-migrate/migrate/blob/master/source/iofs/example_test.go
	migrationFiles, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		log.WithError(err).Error("Could not get migration files")
		return err
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.WithError(err).Error("Could not create db instance")
		return err
	}

	m, err := migrate.NewWithInstance("iofs", migrationFiles, "sqlite3", instance)
	if err != nil {
		log.WithError(err).Error("Could not create migration instance")
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info("Internal DB up do date - no migrations to run")
			return nil
		}

		log.WithError(err).Error("Could not run migrations")
		return err
	}

	log.Info("Run all migrations successfully")
	return nil
}
