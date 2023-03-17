package orchestration

import (
	"github.com/KarnerTh/query-lookout/database"
	"github.com/KarnerTh/query-lookout/notifier"
	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	queryInfra "github.com/KarnerTh/query-lookout/usecase/query/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

func Setup() {
	setupLogger()
	log.Debug("Setup orchestration")

	// database connections
	connectionFactory := database.NewConnectorFactory()
	internalDb, err := connectionFactory.NewConnector("sqlite3://data.db") // TODO: env var?
	if err != nil {
		log.WithError(err).Fatal("Could not initiate internal db")
	}

	watchDb, err := connectionFactory.NewConnector("sqlite3://data.db") // TODO: env var?
	if err != nil {
		log.WithError(err).Fatal("Could not initiate internal db")
	}

	// notifier
	watchResultNotifier := notifier.New[watch.WatchResult]()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(internalDb)
	queryRepo := queryInfra.NewQueryRepo(watchDb)

	// use cases
	watcher := watch.New(watchResultNotifier, queryRepo)
	setupLookout(lookoutRepo, watcher)
}
