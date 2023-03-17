package orchestration

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/notifier"
	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func Setup() {
	log.Debug("Setup orchestration")

	setupLogger()
	db = setupDbConnection()

	// notifier
	watchResultNotifier := notifier.New[watch.WatchResult]()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(db)

	// use cases
	watcher := watch.New(watchResultNotifier)
	setupLookout(lookoutRepo, watcher)
}

func Close() {
	closeDbConection(db)
}
