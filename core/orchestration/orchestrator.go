package orchestration

import (
	"database/sql"

	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func Setup() {
	log.Debug("Setup orchestration")

	setupLogger()
	db = setupDbConnection()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(db)

	// use cases
	setupLookout(lookoutRepo)
}

func Close() {
	closeDbConection(db)
}
