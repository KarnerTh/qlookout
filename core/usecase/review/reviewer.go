package review

import (
	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

type Reviewer interface {
	Start()
	Review(watchResult watch.WatchResult) ([]ReviewResult, error)
}

type reviewer struct {
	watchResultSubscriber watch.WatchResultSubscriber
	reviewResultPublisher ReviewResultPublisher
	reviewRepo            ReviewRepo
}

func New(
	watchResultSubscriber watch.WatchResultSubscriber,
	reviewResultPublisher ReviewResultPublisher,
	reviewRepo ReviewRepo,
) Reviewer {
	return reviewer{
		watchResultSubscriber: watchResultSubscriber,
		reviewResultPublisher: reviewResultPublisher,
		reviewRepo:            reviewRepo,
	}
}

func (r reviewer) Start() {
	watchResultChannel := r.watchResultSubscriber.Subscribe()

	for {
		watchResult := <-watchResultChannel
		reviewResults, _ := r.Review(watchResult)
		for _, result := range reviewResults {
			r.reviewResultPublisher.Publish(result)
		}
	}
}

func (r reviewer) Review(watchResult watch.WatchResult) ([]ReviewResult, error) {
	rules, err := r.reviewRepo.GetRules(watchResult.LookoutId)
	if err != nil {
		log.WithError(err).Errorf("Could not get rules by id")
		return nil, err
	}

	results := make([]ReviewResult, len(rules))
	for i, rule := range rules {
		success, err := validate(watchResult, rule)
		if err != nil {
			// TODO: error part of the review result?
			log.WithError(err).Error("Error in validating rule")
		}

		results[i] = ReviewResult{
			Rule:    rule,
			Success: success,
		}
	}

	return results, nil
}
