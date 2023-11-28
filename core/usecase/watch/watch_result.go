package watch

import (
	"github.com/KarnerTh/qlookout/core/observer"
	"github.com/KarnerTh/qlookout/core/usecase/query"
)

type WatchResultPublisher = observer.Publisher[WatchResult]
type WatchResultSubscriber = observer.Subscriber[WatchResult]

type WatchResult struct {
	LookoutId int
	Result    query.QueryResult
	Error     error
}
