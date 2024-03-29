package graphql

import "github.com/KarnerTh/qlookout/core/usecase/review"

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

func (r ReviewResolver) Rule(args struct{ Id int32 }) (ReviewRuleModel, error) {
	data, err := r.reviewRepo.GetById(int(args.Id))
	if err != nil {
		return nil, err
	}

	return reviewRuleModelResolver{rule: *data}, nil
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

	return reviewRuleModelResolver{rule: *data}, nil
}

func (r ReviewResolver) UpdateRule(args struct {
	Id   int32
	Data reviewRuleUpdateModel
}) (ReviewRuleModel, error) {
	rowIndex := int(*args.Data.RowIndex)
	data, err := r.reviewRepo.Update(
		int(args.Id),
		review.ReviewRuleUpdate{
			ColumnName:   args.Data.ColumnName,
			ColumnType:   (*review.ColumnType)(args.Data.ColumnType),
			RowIndex:     &rowIndex,
			ExactValue:   args.Data.ExactValue,
			LessThan:     args.Data.LessThan,
			GreaterThan:  args.Data.GreaterThan,
			ShouldBeNull: args.Data.ShouldBeNull,
		},
	)

	if err != nil {
		return nil, err
	}

	return reviewRuleModelResolver{rule: *data}, nil
}

func (r ReviewResolver) DeleteRule(args struct{ Id int32 }) (ReviewRuleModel, error) {
	data, err := r.reviewRepo.Delete(int(args.Id))
	if err != nil {
		return nil, err
	}

	return reviewRuleModelResolver{rule: *data}, nil
}
