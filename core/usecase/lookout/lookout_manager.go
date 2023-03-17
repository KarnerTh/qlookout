package lookout

import (
	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

type LookoutManager interface {
	Start()
}

type lookoutManager struct {
	lookoutRepo LookoutRepo
	watcher     watch.Watcher
	cronJobIds  map[int]watch.WatcherId // key is the id of the lookout
}

func NewLookoutService(lookoutRepo LookoutRepo, watcher watch.Watcher) LookoutManager {
	return &lookoutManager{
		lookoutRepo: lookoutRepo,
		watcher:     watcher,
		cronJobIds:  make(map[int]watch.WatcherId),
	}
}

func (l *lookoutManager) Start() {
	log.Debug("Lookout manager started")
	lookouts, err := l.lookoutRepo.Get()
	if err != nil {
		log.WithError(err).Fatal("Could not get lookouts")
	}

	for _, lo := range lookouts {
		id := l.watcher.Watch(watch.WatchConfig{
			LookoutId: lo.Id,
			Name:      lo.Name,
			Query:     lo.Query,
			Cron:      lo.Cron,
		})

		l.cronJobIds[lo.Id] = id
	}
}
