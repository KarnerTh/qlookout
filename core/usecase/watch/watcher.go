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
	cron            *cron.Cron
	resultPublisher WatchResultPublisher
}

func New(resultNotifier WatchResultPublisher) Watcher {
	w := &watcher{
		cron:            cron.New(),
		resultPublisher: resultNotifier,
	}
	w.cron.Start()

	return w
}

func (w watcher) Watch(config WatchConfig) WatcherId {
	job := cronJob[cronJobWatchData]{
		value: cronJobWatchData{
			config:          config,
			resultPublisher: w.resultPublisher,
		},
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

func executeCronJob(job cronJobWatchData) {
	// TODO: execute query
	log.Info("Execute lookout ", job.config.Name)
	job.resultPublisher.Publish(WatchResult{Value: "works from notifier"})
}
