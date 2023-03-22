package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func setupLookout(lookoutRepo lookout.LookoutRepo, watcher watch.Watcher) {
	l := lookout.NewLookoutManager(lookoutRepo, watcher)
	l.Start()
}
