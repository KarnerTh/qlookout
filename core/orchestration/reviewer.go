package orchestration

import (
	"github.com/KarnerTh/qlookout/core/usecase/review"
	"github.com/KarnerTh/qlookout/core/usecase/watch"
)

func setupReviewer(
	watchResultSubscriber watch.WatchResultSubscriber,
	reviewResultPublisher review.ReviewResultPublisher,
	reviewRepo review.ReviewRepo,
) {
	reviewer := review.New(watchResultSubscriber, reviewResultPublisher, reviewRepo)
	go reviewer.Start()
}
