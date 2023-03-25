package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/review"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func setupReviewer(
	watchResultSubscriber watch.WatchResultSubscriber,
	reviewResultPublisher review.ReviewResultPublisher,
	reviewRepo review.ReviewRepo,
) {
	reviewer := review.New(watchResultSubscriber, reviewResultPublisher, reviewRepo)
	go reviewer.Start()
}
