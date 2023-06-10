package lookout

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/core/usecase/watch"
)

type LookoutManager interface {
	Start() error
	Watch(lookoutId int) error
	Remove(lookoutId int) error
	Reload(lookoutId int) error
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

func (l *lookoutManager) Start() error {
	log.Debug("Lookout manager started")
	lookouts, err := l.lookoutRepo.Get()
	if err != nil {
		log.WithError(err).Error("Could not get lookouts")
		return err
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
	return nil
}

func (l *lookoutManager) Watch(lookoutId int) error {
	_, ok := l.watcherIds[lookoutId]
	if ok {
		log.Warnf("Can not add lookout with id %d, because it is already running", lookoutId)
		return fmt.Errorf("Can not add lookout with id %d, because it is already running", lookoutId)
	}

	lookout, err := l.lookoutRepo.GetById(lookoutId)
	if err != nil {
		log.WithError(err).Warnf("Can not add lookout with id %d, because getById failed", lookoutId)
		return err
	}

	id := l.watcher.Watch(watch.WatchConfig{
		LookoutId: lookoutId,
		Name:      lookout.Name,
		Query:     lookout.Query,
		Cron:      lookout.Cron,
	})
	l.watcherIds[lookoutId] = id
	log.Infof("Added watch for lookout with id %d", lookoutId)
	return nil
}

func (l *lookoutManager) Remove(lookoutId int) error {
	watchId, ok := l.watcherIds[lookoutId]
	if !ok {
		log.Errorf("Could not remove lookout with id %d, because it was not found", lookoutId)
		return fmt.Errorf("Could not remove lookout with id %d, because it was not found", lookoutId)
	}

	l.watcher.StopWatching(watchId)
	delete(l.watcherIds, lookoutId)
	log.Infof("Removed watch for lookout with id %d", lookoutId)
	return nil
}

func (l *lookoutManager) Reload(lookoutId int) error {
	err := l.Remove(lookoutId)
	if err != nil {
		return err
	}

	err = l.Watch(lookoutId)
	if err != nil {
		return err
	}

	return nil
}
