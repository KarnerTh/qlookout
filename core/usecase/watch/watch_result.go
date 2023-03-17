package watch

import "github.com/KarnerTh/query-lookout/notifier"

type WatchResultPublisher = notifier.Publisher[WatchResult]
type WatchResultSubscriber = notifier.Subscriber[WatchResult]

type WatchResult struct {
  Value string
}
