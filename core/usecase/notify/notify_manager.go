package notify

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/review"
)

type NotifyManager interface {
	Start()
	Notify(reviewResult review.ReviewResult)
}

type notifyManager struct {
	reviewResultSubscriber review.ReviewResultSubscriber
	lookoutService         lookout.LookoutService
	localNotifier          Notifier
	mailNotifier           Notifier
}

func New(reviewResultSubscriber review.ReviewResultSubscriber, lookoutService lookout.LookoutService, localNotifier Notifier, mailNotifier Notifier) NotifyManager {
	return notifyManager{
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

		if lookout.NotifyLocal {
			_ = n.localNotifier.Send(Notification{Title: fmt.Sprintf("NOK: %s", lookout.Name), Description: "rule not successfull"})
		}
		if lookout.NotifyMail {
			_ = n.mailNotifier.Send(Notification{Title: fmt.Sprintf("NOK: %s", lookout.Name), Description: "rule not successfull"})
		}
	}
}
