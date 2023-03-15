package lookout

import (
	"github.com/KarnerTh/query-lookout/cronjob"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

type LookoutService interface {
	Start()
}

type lookoutService struct {
	cron        cron.Cron
	lookoutRepo LookoutRepo
}

func NewLookoutService(lookoutRepo LookoutRepo) LookoutService {
	return &lookoutService{
		cron:        *cron.New(),
		lookoutRepo: lookoutRepo,
	}
}

func (l *lookoutService) Start() {
	log.Debug("Lookout started")
	lookouts, err := l.lookoutRepo.Get()
	if err != nil {
		log.WithError(err).Fatal("Could not get lookouts")
	}

	for _, lo := range lookouts {
		job := cronjob.CronJob[string]{
			Value: lo.Query,
			Execute: func(value string) {
				log.Info("nice ", value)
			},
		}

		err := l.cron.AddJob(lo.Cron, job)
		if err != nil {
			log.WithError(err).Fatal("Could not start cron schedule")
		}
	}
	l.cron.Start()
}
