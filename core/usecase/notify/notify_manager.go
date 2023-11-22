package notify

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/KarnerTh/query-lookout/core/config"
	"github.com/KarnerTh/query-lookout/core/usecase/lookout"
	"github.com/KarnerTh/query-lookout/core/usecase/review"
)

type NotifyManager interface {
	Start()
	Notify(reviewResult review.ReviewResult)
}

type notifyManager struct {
	config                 config.Config
	reviewResultSubscriber review.ReviewResultSubscriber
	notificationPublisher  NotificationPublisher
	lookoutRepo            lookout.LookoutRepo
	localNotifier          Notifier
	mailNotifier           Notifier
}

func New(config config.Config, reviewResultSubscriber review.ReviewResultSubscriber, notificationPublisher NotificationPublisher, lookoutRepo lookout.LookoutRepo, localNotifier Notifier, mailNotifier Notifier) NotifyManager {
	return notifyManager{
		config:                 config,
		reviewResultSubscriber: reviewResultSubscriber,
		notificationPublisher:  notificationPublisher,
		lookoutRepo:            lookoutRepo,
		localNotifier:          localNotifier,
		mailNotifier:           mailNotifier,
	}
}

func (n notifyManager) Start() {
	reviewResultChannel := n.reviewResultSubscriber.Subscribe()

	for {
		reviewResult := <-reviewResultChannel
		n.Notify(reviewResult)
	}
}

func (n notifyManager) Notify(reviewResult review.ReviewResult) {
	if !reviewResult.Result.IsValid {
		lookout, err := n.lookoutRepo.GetById(reviewResult.LookoutId)
		if err != nil {
			slog.Error("Could not get lookout config", slog.Any("error", err))
			return
		}

		var notification Notification

		if reviewResult.Error == nil {
			notification = Notification{
				LookoutId:   lookout.Id,
				RuleId:      reviewResult.Rule.Id,
				Title:       fmt.Sprintf("NOK: %s", lookout.Name),
				Description: reviewResult.Result.Description,
				Timestamp:   time.Now(),
			}
		} else {
			notification = Notification{
				LookoutId:   lookout.Id,
				Title:       fmt.Sprintf("Error: %s", lookout.Name),
				Description: reviewResult.Error.Error(),
				Timestamp:   time.Now(),
			}
		}

		n.notificationPublisher.Publish(notification)

		if lookout.NotifyLocal {
			err = n.localNotifier.Send(notification)
			if err != nil {
				slog.Error("Could not send local notification", slog.Any("error", err))
			}
		}
		if lookout.NotifyMail {
			err = n.mailNotifier.Send(notification)
			if err != nil {
				slog.Error("Could not send mail notification", slog.Any("error", err))
			}
		}
	}
}
