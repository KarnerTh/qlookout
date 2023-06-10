package watch

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/core/usecase/query"
)

type WatcherId = cron.EntryID

//go:generate mockery --name Watcher
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
		log.WithError(err).Fatal("Could not start cron job - please check lookout configuration")
	}

	return id
}

func (w watcher) StopWatching(id WatcherId) {
	w.cron.Remove(id)
}

func executeCronJob(job cronJobWatchData) {
	log.Infof("Execute lookout %s", job.config.Name)
	result, err := job.queryRepo.Query(job.config.Query)
	if err != nil {
		log.WithError(err).Error("Error quering job")
		job.resultPublisher.Publish(WatchResult{LookoutId: job.config.LookoutId, Error: err})
		return
	}

	job.resultPublisher.Publish(WatchResult{LookoutId: job.config.LookoutId, Result: result})
}
