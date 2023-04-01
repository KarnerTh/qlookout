package infrastructure

import (
	"github.com/KarnerTh/query-lookout/usecase/notify"
	"github.com/gen2brain/beeep"
	log "github.com/sirupsen/logrus"
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
