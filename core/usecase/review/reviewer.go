package review

import (
	"log/slog"

	"github.com/KarnerTh/query-lookout/core/usecase/watch"
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
	rules, err := r.reviewRepo.GetForLookout(watchResult.LookoutId)
	if err != nil {
		slog.Error("Could not get rules by id", slog.Any("error", err))
		return nil, err
	}

	if watchResult.Error != nil {
		return []ReviewResult{{LookoutId: watchResult.LookoutId, Result: ValidationResult{IsValid: false}, Error: watchResult.Error}}, nil
	}

	results := make([]ReviewResult, len(rules))
	for i, rule := range rules {
		validationResult, err := validate(watchResult, rule)
		if err != nil {
			slog.Error("Error in validating rule", slog.Any("error", err))
		}

		results[i] = ReviewResult{
			LookoutId: watchResult.LookoutId,
			Rule:      rule,
			Result:    validationResult,
			Error:     err,
		}
	}

	return results, nil
}
