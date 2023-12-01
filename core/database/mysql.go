package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func openMysql(driverName DbType, dataSource string) (*sql.DB, error) {
	connectionString := strings.ReplaceAll(dataSource, "mysql://", "")
	db, err := sql.Open(driverName.String(), connectionString)
	if err != nil {
		slog.Error(fmt.Sprintf("Could not open database %s with connection string %s", driverName, dataSource), slog.Any("error", err))
		return nil, err
	}

	return db, nil
}
