package watch

import "github.com/KarnerTh/query-lookout/usecase/query"

type cronJobWatchData struct {
	config          WatchConfig
	resultPublisher WatchResultPublisher
	queryRepo       query.QueryRepo
}
