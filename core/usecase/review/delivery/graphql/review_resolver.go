package graphql

import "github.com/KarnerTh/query-lookout/usecase/review"

type ReviewResolver struct {
	reviewRepo review.ReviewRepo
}

func NewReviewResolver(reviewRepo review.ReviewRepo) ReviewResolver {
	return ReviewResolver{
		reviewRepo: reviewRepo,
	}
}

func (r ReviewResolver) Rules(args struct{ LookoutId int32 }) ([]ReviewRuleModel, error) {
	rules, err := r.reviewRepo.GetRules(int(args.LookoutId))
	if err != nil {
		return []ReviewRuleModel{}, err
	}

	models := make([]ReviewRuleModel, len(rules))
	for i, value := range rules {
		models[i] = reviewRuleModelResolver{rule: value}
	}

	return models, nil
}
