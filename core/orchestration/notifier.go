package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/notify"
	"github.com/KarnerTh/query-lookout/usecase/review"
)

func setupNotifier(reviewResultSubscriber review.ReviewResultSubscriber, lookoutService lookout.LookoutService) {
	notifier := notify.New(reviewResultSubscriber, lookoutService)
	go notifier.Start()
}
