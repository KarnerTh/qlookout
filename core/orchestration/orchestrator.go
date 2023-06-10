package orchestration

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/core/database"
	"github.com/KarnerTh/query-lookout/core/observer"
	lookoutInfra "github.com/KarnerTh/query-lookout/core/usecase/lookout/infrastructure"
	"github.com/KarnerTh/query-lookout/core/usecase/notify"
	queryInfra "github.com/KarnerTh/query-lookout/core/usecase/query/infrastructure"
	"github.com/KarnerTh/query-lookout/core/usecase/review"
	reviewInfra "github.com/KarnerTh/query-lookout/core/usecase/review/infrastructure"
	"github.com/KarnerTh/query-lookout/core/usecase/watch"
)

func Setup() {
	config := setupConfig()
	setupLogger(config)
	log.Debug("Setup orchestration")

	// database connections
	connectionFactory := database.NewConnectorFactory()
	internalDb, err := connectionFactory.NewConnector("sqlite3://data.db")
	defer func() { _ = internalDb.Close() }()
	if err != nil {
		log.WithError(err).Fatal("Could not initiate internal db")
	}

	// enable sqlite foreign key support
	_, err = internalDb.Exec("PRAGMA foreign_keys=ON")
	if err != nil {
		log.WithError(err).Fatal("Could not enable foreign key support for internal db")
	}

	err = database.MigrateInternalSqliteDb(internalDb)
	if err != nil {
		log.WithError(err).Fatal("Shutting down due to failing migrations")
	}

	watchDb, err := connectionFactory.NewConnector(config.DataSource())
	defer func() { _ = internalDb.Close() }()
	if err != nil {
		log.WithError(err).Fatal("Could not initiate connection to data source")
	}

	// notifier
	watchResultObserver := observer.New[watch.WatchResult]()
	reviewResultObserver := observer.New[review.ReviewResult]()
	notificationObserver := observer.New[notify.Notification]()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(internalDb)
	queryRepo := queryInfra.NewQueryRepo(watchDb)
	reviewRepo := reviewInfra.NewReviewRepo(internalDb)

	// use cases
	watcher := watch.New(watchResultObserver, queryRepo)
	lookoutManager := setupLookout(lookoutRepo, watcher)
	setupReviewer(watchResultObserver, reviewResultObserver, reviewRepo)
	setupNotifier(config, reviewResultObserver, notificationObserver, lookoutRepo)

	// delivery
	setupDelivery(lookoutManager, lookoutRepo, reviewRepo, notificationObserver)
}
