package orchestration

import (
	"fmt"
	"log/slog"

	"github.com/KarnerTh/qlookout/core/database"
	"github.com/KarnerTh/qlookout/core/observer"
	lookoutInfra "github.com/KarnerTh/qlookout/core/usecase/lookout/infrastructure"
	"github.com/KarnerTh/qlookout/core/usecase/notify"
	queryInfra "github.com/KarnerTh/qlookout/core/usecase/query/infrastructure"
	"github.com/KarnerTh/qlookout/core/usecase/review"
	reviewInfra "github.com/KarnerTh/qlookout/core/usecase/review/infrastructure"
	"github.com/KarnerTh/qlookout/core/usecase/watch"
)

func Setup() {
	flags := parseFlags()
	config := setupConfig(flags.ConfigPath)
	setupLogger(config)
	slog.Debug("Setup orchestration")

	// database connections
	connectionFactory := database.NewConnectorFactory()
	internalDb, err := connectionFactory.NewConnector(fmt.Sprintf("sqlite3://%s", config.DatabaseFile()))
	defer func() { _ = internalDb.Close() }()
	if err != nil {
		slog.Error("Could not initiate internal db", slog.Any("error", err))
		return
	}

	// enable sqlite foreign key support
	_, err = internalDb.Exec("PRAGMA foreign_keys=ON")
	if err != nil {
		slog.Error("Could not enable foreign key support for internal db", slog.Any("error", err))
		return
	}

	err = database.MigrateInternalSqliteDb(internalDb)
	if err != nil {
		slog.Error("Shutting down due to failing migrations", slog.Any("error", err))
		return
	}

	watchDb, err := connectionFactory.NewConnector(config.DataSource())
	defer func() { _ = internalDb.Close() }()
	if err != nil {
		slog.Error("Could not initiate connection to data source", slog.Any("error", err))
		return
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
