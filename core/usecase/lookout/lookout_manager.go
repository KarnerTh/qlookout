package lookout

import (
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/usecase/watch"
)

type LookoutManager interface {
	Start()
	Watch(lookoutId int)
	Remove(lookoutId int)
	Reload(lookoutId int)
}

type lookoutManager struct {
	lookoutRepo LookoutRepo
	watcher     watch.Watcher
	watcherIds  map[int]watch.WatcherId // key is the id of the lookout
}

func NewLookoutManager(lookoutRepo LookoutRepo, watcher watch.Watcher) LookoutManager {
	return &lookoutManager{
		lookoutRepo: lookoutRepo,
		watcher:     watcher,
		watcherIds:  make(map[int]watch.WatcherId),
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
		l.watcherIds[lo.Id] = id
	}
	log.Info("All lookouts started successfully")
}

func (l *lookoutManager) Watch(lookoutId int) {
	_, ok := l.watcherIds[lookoutId]
	if ok {
		log.Warnf("Can not add lookout with id %d, because it is already running", lookoutId)
		return
	}

	lookout, err := l.lookoutRepo.GetById(lookoutId)
	if err != nil {
		log.WithError(err).Warnf("Can not add lookout with id %d, because getById failed", lookoutId)
		return
	}

	id := l.watcher.Watch(watch.WatchConfig{
		LookoutId: lookoutId,
		Name:      lookout.Name,
		Query:     lookout.Query,
		Cron:      lookout.Cron,
	})
	l.watcherIds[lookoutId] = id
	log.Infof("Added watch for lookout with id %d", lookoutId)
}

func (l *lookoutManager) Remove(lookoutId int) {
	watchId, ok := l.watcherIds[lookoutId]
	if !ok {
		log.Errorf("Could not remove lookout with id %d, because it was not found", lookoutId)
	}

	l.watcher.StopWatching(watchId)
	delete(l.watcherIds, lookoutId)
	log.Infof("Removed watch for lookout with id %d", lookoutId)
}

func (l *lookoutManager) Reload(lookoutId int) {
	l.Remove(lookoutId)
	l.Watch(lookoutId)
}
