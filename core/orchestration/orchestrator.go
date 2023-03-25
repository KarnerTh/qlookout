package orchestration

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/database"
	"github.com/KarnerTh/query-lookout/notifier"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	lookoutInfra "github.com/KarnerTh/query-lookout/usecase/lookout/infrastructure"
	queryInfra "github.com/KarnerTh/query-lookout/usecase/query/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/review"
	reviewInfra "github.com/KarnerTh/query-lookout/usecase/review/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/watch"
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
	reviewResultNotifier := notifier.New[review.ReviewResult]()

	// repos
	lookoutRepo := lookoutInfra.NewLookoutRepo(internalDb)
	queryRepo := queryInfra.NewQueryRepo(watchDb)
	reviewRepo := reviewInfra.NewReviewRepo(internalDb)

	// services
	lookoutService := lookout.NewLookoutService(lookoutRepo)

	// use cases
	watcher := watch.New(watchResultNotifier, queryRepo)
	setupLookout(lookoutService, watcher)
	setupReviewer(watchResultNotifier, reviewResultNotifier, reviewRepo)
}
