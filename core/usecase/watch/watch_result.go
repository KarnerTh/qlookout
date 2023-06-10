package watch

import (
	"github.com/KarnerTh/query-lookout/core/observer"
	"github.com/KarnerTh/query-lookout/core/usecase/query"
)

type WatchResultPublisher = observer.Publisher[WatchResult]
type WatchResultSubscriber = observer.Subscriber[WatchResult]

type WatchResult struct {
	LookoutId int
	Result    query.QueryResult
	Error     error
}
