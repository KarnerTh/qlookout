package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/review"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func setupReviewer(watchResultSubscriber watch.WatchResultSubscriber, reviewRepo review.ReviewRepo) {
	reviewer := review.New(watchResultSubscriber, reviewRepo)
	go reviewer.Start()
}
