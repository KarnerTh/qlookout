package infrastructure

import (
	"github.com/gen2brain/beeep"
	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/core/usecase/notify"
)

type localNotifier struct{}

func NewLocalNotifier() notify.Notifier {
	return localNotifier{}
}

func (n localNotifier) Send(value notify.Notification) error {
	err := beeep.Notify(value.Title, value.Description, "assets/information.png")
	if err != nil {
		log.WithError(err).Error("Could not create local notifiation")
	}

	return err
}
