package orchestration

import (
	"github.com/KarnerTh/query-lookout/core/config"
	"github.com/KarnerTh/query-lookout/core/usecase/lookout"
	"github.com/KarnerTh/query-lookout/core/usecase/notify"
	notifyInfra "github.com/KarnerTh/query-lookout/core/usecase/notify/infrastructure"
	"github.com/KarnerTh/query-lookout/core/usecase/review"
)

func setupNotifier(config config.Config, reviewResultSubscriber review.ReviewResultSubscriber, notificationPublisher notify.NotificationPublisher, lookoutRepo lookout.LookoutRepo) {
	localNotifier := notifyInfra.NewLocalNotifier()
	mailNotifier := notifyInfra.NewMailNotifier(config)

	notifier := notify.New(config, reviewResultSubscriber, notificationPublisher, lookoutRepo, localNotifier, mailNotifier)
	go notifier.Start()
}
