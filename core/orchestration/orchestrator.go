package orchestration

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/database"
	"github.com/KarnerTh/query-lookout/observer"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	queryInfra "github.com/KarnerTh/query-lookout/usecase/query/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/review"
	reviewInfra "github.com/KarnerTh/query-lookout/usecase/review/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func Setup() {
	config := setupConfig()
	setupLogger()
	log.Debug("Setup orchestration")

	// database connections
	connectionFactory := database.NewConnectorFactory()
	internalDb, err := connectionFactory.NewConnector("sqlite3://data.db")
	if err != nil {
		log.WithError(err).Fatal("Could not initiate internal db")
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

	// services
	lookoutService := lookout.NewLookoutService(lookoutRepo)

	// use cases
	watcher := watch.New(watchResultObserver, queryRepo)
	setupLookout(lookoutService, watcher)
	setupReviewer(watchResultObserver, reviewResultObserver, reviewRepo)
	setupNotifier(config, reviewResultObserver, lookoutService)
}
