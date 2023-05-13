package orchestration

import (
	"github.com/KarnerTh/query-lookout/config"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/notify"
	notifyInfra "github.com/KarnerTh/query-lookout/usecase/notify/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/review"
)

func setupNotifier(config config.Config, reviewResultSubscriber review.ReviewResultSubscriber, notificationPublisher notify.NotificationPublisher, lookoutRepo lookout.LookoutRepo) {
	localNotifier := notifyInfra.NewLocalNotifier()
	mailNotifier := notifyInfra.NewMailNotifier(config.MailFromAddress(), config.MailToAddress(), config.MailSmtpHost(), config.MailSmtpPort())

	notifier := notify.New(config, reviewResultSubscriber, notificationPublisher, lookoutRepo, localNotifier, mailNotifier)
	go notifier.Start()
}
