package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

func openSqlite(driverName DbType, dataSource string) (*sql.DB, error) {
	connectionString := strings.ReplaceAll(dataSource, "sqlite3://", "file:")
	db, err := sql.Open(driverName.String(), connectionString)
	if err != nil {
		slog.Error(fmt.Sprintf("Could not open database %s with connection string %s", driverName, dataSource), slog.Any("error", err))
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
		slog.Error("Could not get migration files", slog.Any("error", err))
		return err
	}

	instance, err := WithSqliteMigrateInstance(db, &SqliteMigrateConfig{})
	if err != nil {
		slog.Error("Could not create db instance", slog.Any("error", err))
		return err
	}

	m, err := migrate.NewWithInstance("iofs", migrationFiles, "sqlite", instance)
	if err != nil {
		slog.Error("Could not create migration instance", slog.Any("error", err))
		return err
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("Internal DB up do date - no migrations to run")
			return nil
		}

		slog.Error("Could not run migrations", slog.Any("error", err))
		return err
	}

	slog.Info("Run all migrations successfully")
	return nil
}
