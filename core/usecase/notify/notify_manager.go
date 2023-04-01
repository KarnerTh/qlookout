package notify

import (
	"fmt"

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
	lookoutService         lookout.LookoutService
	localNotifier          Notifier
	mailNotifier           Notifier
}

func New(config config.Config, reviewResultSubscriber review.ReviewResultSubscriber, lookoutService lookout.LookoutService, localNotifier Notifier, mailNotifier Notifier) NotifyManager {
	return notifyManager{
		config:                 config,
		reviewResultSubscriber: reviewResultSubscriber,
		lookoutService:         lookoutService,
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
	if !reviewResult.Success {
		lookout, err := n.lookoutService.GetConfig(reviewResult.Rule.LookoutId)
		if err != nil {
			log.WithError(err).Error("Could not get lookout config")
		}

		notification := Notification{
			Title:       fmt.Sprintf("NOK: %s", lookout.Name),
			Description: "rule not successfull",
			DeepLink:    n.config.BaseUrl(), // TODO: add deeplink parameter
		}

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
