package watch

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type WatcherId = cron.EntryID

type Watcher interface {
	Watch(config WatchConfig) WatcherId
	StopWatching(id WatcherId)
}

type watcher struct {
	cron *cron.Cron
}

func New() Watcher {
	w := &watcher{
		cron: cron.New(),
	}
	w.cron.Start()

	return w
}

func (w watcher) Watch(config WatchConfig) WatcherId {
	job := cronJob[WatchConfig]{
		value:   config,
		execute: executeCronJob,
	}

	id, err := w.cron.AddJob(config.Cron, job)
	if err != nil {
		log.WithError(err).Fatal("Could not start cron schedule")
	}

	return id
}
func (w watcher) StopWatching(id WatcherId) {
	w.cron.Remove(id)
}

func executeCronJob(job WatchConfig) {
	log.Info("Execute lookout ", job.Name)
	// TODO: execute query
}
