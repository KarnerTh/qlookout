package orchestration

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

// TODO: env variable?
const dbPath = "file:data.db"

func setupDbConnection() *sql.DB {
	log.Debug("Connect to internal db")
	con, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.WithError(err).Fatal("Could not open internal")
	}

	err = con.Ping()
	if err != nil {
		log.WithError(err).Fatal("Could not verify db connection")
	}

	log.Debug("Connected to internal db successfully")
	return con
}

func closeDbConection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.WithError(err).Fatal("Could not close db")
	}
}
