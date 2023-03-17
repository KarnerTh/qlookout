package watch

import (
	"github.com/KarnerTh/query-lookout/usecase/query"
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
	queryRepo       query.QueryRepo
}

func New(resultNotifier WatchResultPublisher, queryRepo query.QueryRepo) Watcher {
	w := &watcher{
		cron:            cron.New(),
		resultPublisher: resultNotifier,
		queryRepo:       queryRepo,
	}
	w.cron.Start()

	return w
}

func (w watcher) Watch(config WatchConfig) WatcherId {
	job := cronJob[cronJobWatchData]{
		value: cronJobWatchData{
			config:          config,
			resultPublisher: w.resultPublisher,
			queryRepo:       w.queryRepo,
		},
		execute: executeCronJob,
	}

	id, err := w.cron.AddJob(config.Cron, job)
	if err != nil {
		log.WithError(err).Fatal("Could not cron job - please check lookout configuration")
	}

	return id
}

func (w watcher) StopWatching(id WatcherId) {
	w.cron.Remove(id)
}

func executeCronJob(job cronJobWatchData) {
	// TODO: execute query
	log.Info("Execute lookout ", job.config.Name)
	result, err := job.queryRepo.Query(job.config.Query)
	if err != nil {
		log.WithError(err).Error("Error quering job")
	}

	log.Infof("Result: %+v\n", result)
	job.resultPublisher.Publish(WatchResult{Value: "works from notifier"})
}
