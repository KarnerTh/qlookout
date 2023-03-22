package watch

import (
	"github.com/KarnerTh/query-lookout/notifier"
	"github.com/KarnerTh/query-lookout/usecase/query"
)

type WatchResultPublisher = notifier.Publisher[WatchResult]
type WatchResultSubscriber = notifier.Subscriber[WatchResult]

type WatchResult struct {
	LookoutId int
	Result    query.QueryResult
}
