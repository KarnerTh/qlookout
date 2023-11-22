package infrastructure

import (
	"log/slog"

	"github.com/gen2brain/beeep"

	"github.com/KarnerTh/query-lookout/core/usecase/notify"
)

type localNotifier struct{}

func NewLocalNotifier() notify.Notifier {
	return localNotifier{}
}

func (n localNotifier) Send(value notify.Notification) error {
	err := beeep.Notify(value.Title, value.Description, "assets/information.png")
	if err != nil {
		slog.Error("Could not create local notifiation", slog.Any("error", err))
	}

	return err
}
