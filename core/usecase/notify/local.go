package notify

import (
	"github.com/gen2brain/beeep"
	log "github.com/sirupsen/logrus"
)

func sendLocalNotification(value Notification) {
	err := beeep.Notify(value.Title, value.Description, "assets/information.png")
	if err != nil {
		log.WithError(err).Error("Could not create local notifiation")
	}
}
