package watch

import (
	"github.com/KarnerTh/query-lookout/observer"
	"github.com/KarnerTh/query-lookout/usecase/query"
)

type WatchResultPublisher = observer.Publisher[WatchResult]
type WatchResultSubscriber = observer.Subscriber[WatchResult]

type WatchResult struct {
	LookoutId int
	Result    query.QueryResult
}
