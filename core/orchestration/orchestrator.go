package orchestration

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/database"
	"github.com/KarnerTh/query-lookout/observer"
	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	queryInfra "github.com/KarnerTh/query-lookout/usecase/query/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/review"
	reviewInfra "github.com/KarnerTh/query-lookout/usecase/review/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func Setup() {
	config := setupConfig()
	setupLogger(config)
	log.Debug("Setup orchestration")

	// database connections
	connectionFactory := database.NewConnectorFactory()
	internalDb, err := connectionFactory.NewConnector("sqlite3://data.db")
	if err != nil {
		log.WithError(err).Fatal("Could not initiate internal db")
	}

	err = database.MigrateInternalSqliteDb(internalDb)
	if err != nil {
		log.WithError(err).Fatal("Shutting down due to failing migrations")
	}

	watchDb, err := connectionFactory.NewConnector(config.DataSource())
	if err != nil {
		log.WithError(err).Fatal("Could not initiate connection to data source")
	}

	// notifier
	watchResultObserver := observer.New[watch.WatchResult]()
	reviewResultObserver := observer.New[review.ReviewResult]()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(internalDb)
	queryRepo := queryInfra.NewQueryRepo(watchDb)
	reviewRepo := reviewInfra.NewReviewRepo(internalDb)

	// use cases
	watcher := watch.New(watchResultObserver, queryRepo)
	lookoutManager := setupLookout(lookoutRepo, watcher)
	setupReviewer(watchResultObserver, reviewResultObserver, reviewRepo)
	setupNotifier(config, reviewResultObserver, lookoutRepo)

	// delivery
	setupDelivery(lookoutManager, lookoutRepo, reviewRepo)
}
