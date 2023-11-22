package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func openPostgres(driverName DbType, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driverName.String(), dataSource)
	if err != nil {
		slog.Error(fmt.Sprintf("Could not open database %s with connection string %s", driverName, dataSource), slog.Any("error", err))
		return nil, err
	}

	return db, nil
}
