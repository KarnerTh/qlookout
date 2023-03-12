package lookout

import (
	"fmt"

	"github.com/robfig/cron"

	"github.com/KarnerTh/query-lookout/domain"
)

type Lookout interface {
	Start()
}

type lookout struct {
	cron cron.Cron
}

func New(config domain.LookoutConfig) Lookout {
	return &lookout{cron: *cron.New()}
}

func (l *lookout) Start() {
	fmt.Println("Lookout started")

	// TODO: add all cron jobs according to config
	l.cron.AddFunc("@every 1s", func() { fmt.Println("Every second") })

	l.cron.Start()
}
