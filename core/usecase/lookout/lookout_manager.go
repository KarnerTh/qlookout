package lookout

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/usecase/watch"
)

type LookoutManager interface {
	Start()
}

type lookoutManager struct {
	lookoutService LookoutService
	watcher        watch.Watcher
	cronJobIds     map[int]watch.WatcherId // key is the id of the lookout
}

func NewLookoutManager(lookoutService LookoutService, watcher watch.Watcher) LookoutManager {
	return &lookoutManager{
		lookoutService: lookoutService,
		watcher:        watcher,
		cronJobIds:     make(map[int]watch.WatcherId),
	}
}

func (l *lookoutManager) Start() {
	log.Debug("Lookout manager started")
	lookouts, err := l.lookoutService.GetConfigs()
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
	log.Debug("All lookouts started successfully")
}
