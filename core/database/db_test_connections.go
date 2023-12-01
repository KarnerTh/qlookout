package database

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var directory string

type connectionData struct {
	DbType           DbType
	ConnectionString string
}

func GetTestDbConnections() []connectionData {
	directory, err := os.Getwd()
	if err != nil {
		slog.Error("error in getting current workdir for tests", slog.Any("error", err))
	}

	basePath := strings.Split(directory, "/core")[0]

	return []connectionData{
		{
			DbType:           Postgres,
			ConnectionString: "postgresql://user:password@localhost:5432/qlookout_test",
		},
		{
			DbType:           MySql,
			ConnectionString: "mysql://user:password@tcp(127.0.0.1:3306)/qlookout_test",
		},
		{
			DbType:           MsSql,
			ConnectionString: "sqlserver://sa:securePassword1!@localhost:1433?database=qlookout_test",
		},
		{
			DbType:           Sqlite3,
			ConnectionString: fmt.Sprintf("sqlite3://%s/core/qlookout_test.db", basePath),
		},
	}
}
