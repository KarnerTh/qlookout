package watch

import "github.com/KarnerTh/qlookout/core/usecase/query"

type cronJobWatchData struct {
	config          WatchConfig
	resultPublisher WatchResultPublisher
	queryRepo       query.QueryRepo
}
