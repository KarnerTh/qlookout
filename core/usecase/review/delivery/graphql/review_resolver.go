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
	rules, err := r.reviewRepo.GetForLookout(int(args.LookoutId))
	if err != nil {
		return []ReviewRuleModel{}, err
	}

	models := make([]ReviewRuleModel, len(rules))
	for i, value := range rules {
		models[i] = reviewRuleModelResolver{rule: value}
	}

	return models, nil
}

func (r ReviewResolver) CreateRule(args struct{ Data reviewRuleCreateModel }) (ReviewRuleModel, error) {
	data, err := r.reviewRepo.Create(review.ReviewRuleCreate{
		LookoutId:    int(args.Data.LookoutId),
		ColumnName:   args.Data.ColumnName,
		ColumnType:   review.ColumnType(args.Data.ColumnType),
		RowIndex:     int(args.Data.RowIndex),
		ExactValue:   args.Data.ExactValue,
		ShouldBeNull: args.Data.ShouldBeNull,
		LessThan:     args.Data.LessThan,
		GreaterThan:  args.Data.GreaterThan,
	})

	if err != nil {
		return nil, err
	}

	return reviewRuleModelResolver{rule: data}, nil
}
