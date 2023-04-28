package orchestration

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func setupLookout(lookoutRepo lookout.LookoutRepo, watcher watch.Watcher) lookout.LookoutManager {
	l := lookout.NewLookoutManager(lookoutRepo, watcher)

	go func() {
		err := l.Start()
		if err != nil {
			panic(err)
		}
	}()
	return l
}
