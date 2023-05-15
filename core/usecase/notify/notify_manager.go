package notify

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/config"
	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/review"
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
			log.WithError(err).Error("Could not get lookout config")
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
				log.WithError(err).Error("Could not send local notification")
			}
		}
		if lookout.NotifyMail {
			err = n.mailNotifier.Send(notification)
			if err != nil {
				log.WithError(err).Error("Could not send mail notification")
			}
		}
	}
}
