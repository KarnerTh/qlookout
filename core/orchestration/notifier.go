package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/notify"
	notifyInfra "github.com/KarnerTh/query-lookout/usecase/notify/infrastructure"
	"github.com/KarnerTh/query-lookout/usecase/review"
)

func setupNotifier(reviewResultSubscriber review.ReviewResultSubscriber, lookoutService lookout.LookoutService) {
	localNotifier := notifyInfra.NewLocalNotifier()
	mailNotifier := notifyInfra.NewMailNotifier("testAddress@tka.com", "localhost", "1025")

	notifier := notify.New(reviewResultSubscriber, lookoutService, localNotifier, mailNotifier)
	go notifier.Start()
}
