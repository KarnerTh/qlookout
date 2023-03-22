package review

import (
	"fmt"

	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

type Reviewer interface {
	Start()
}

type reviewer struct {
	watchResultSubscriber watch.WatchResultSubscriber
	reviewRepo            ReviewRepo
}

func New(watchResultSubscriber watch.WatchResultSubscriber, reviewRepo ReviewRepo) Reviewer {
	return &reviewer{
		watchResultSubscriber: watchResultSubscriber,
		reviewRepo:            reviewRepo,
	}
}

func (r reviewer) Start() {
	watchResultChannel := r.watchResultSubscriber.Subscribe()

	for {
		watchResult := <-watchResultChannel

		reviewResults := r.Review(watchResult)
		for _, result := range reviewResults {
			if !result.Success {
				log.Error("heyyy")
			}
		}
	}
}

func (r reviewer) Review(watchResult watch.WatchResult) []ReviewResult {
	rules, err := r.reviewRepo.GetRules(watchResult.LookoutId)
	if err != nil {
		log.WithError(err).Errorf("Could not get rules by id")
	}

	results := make([]ReviewResult, len(rules))
	for i, rule := range rules {
		expectedValue := rule.ExactValue
		actualValue := fmt.Sprint(watchResult.Result.Rows[rule.RowIndex][rule.ColumnName])

		results[i] = ReviewResult{
			Rule:    rule,
			Success: actualValue == expectedValue,
		}
	}

	return results
}
