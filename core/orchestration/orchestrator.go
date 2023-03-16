package orchestration

import (
	"database/sql"

	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
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
	watcher := watch.New()
	setupLookout(lookoutRepo, watcher)
}

func Close() {
	closeDbConection(db)
}
