package notify

import (
	"fmt"

	"github.com/KarnerTh/query-lookout/usecase/lookout"
	"github.com/KarnerTh/query-lookout/usecase/review"
	log "github.com/sirupsen/logrus"
)

type Notifier interface {
	Start()
	Notify(reviewResult review.ReviewResult)
}

type notifier struct {
	reviewResultSubscriber review.ReviewResultSubscriber
	lookoutService         lookout.LookoutService
}

func New(reviewResultSubscriber review.ReviewResultSubscriber, lookoutService lookout.LookoutService) Notifier {
	return notifier{
		reviewResultSubscriber: reviewResultSubscriber,
		lookoutService:         lookoutService,
	}
}

func (n notifier) Start() {
	reviewResultChannel := n.reviewResultSubscriber.Subscribe()

	for {
		reviewResult := <-reviewResultChannel
		n.Notify(reviewResult)
	}
}

func (n notifier) Notify(reviewResult review.ReviewResult) {
	if !reviewResult.Success {
		lookout, err := n.lookoutService.GetConfig(reviewResult.Rule.LookoutId)
		if err != nil {
			log.WithError(err).Error("Could not get lookout config")
		}

		if lookout.NotifyLocal {
			sendLocalNotification(Notification{Title: fmt.Sprintf("NOK: %s", lookout.Name), Description: "rule not successfull"})
		}
		if lookout.NotifyMail {
			sendMailNotification(Notification{Title: fmt.Sprintf("NOK: %s", lookout.Name), Description: "rule not successfull"})
		}
	}
}
