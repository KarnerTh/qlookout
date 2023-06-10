package watch

import "github.com/KarnerTh/query-lookout/core/usecase/query"

type cronJobWatchData struct {
	config          WatchConfig
	resultPublisher WatchResultPublisher
	queryRepo       query.QueryRepo
}
