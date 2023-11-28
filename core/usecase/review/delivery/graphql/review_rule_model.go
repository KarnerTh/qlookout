package graphql

import "github.com/KarnerTh/qlookout/core/usecase/review"

type ReviewRuleModel interface {
	Id() int32
	LookoutId() int32
	ColumnName() string
	ColumnType() string
	RowIndex() int32
	ExactValue() *string
	GreaterThan() *string
	LessThan() *string
	ShouldBeNull() bool
}

type reviewRuleModelResolver struct {
	rule review.ReviewRule
}

func (r reviewRuleModelResolver) Id() int32 {
	return int32(r.rule.Id)
}

func (r reviewRuleModelResolver) LookoutId() int32 {
	return int32(r.rule.LookoutId)
}

func (r reviewRuleModelResolver) ColumnName() string {
	return r.rule.ColumnName
}

func (r reviewRuleModelResolver) ColumnType() string {
	return string(r.rule.ColumnType)
}

func (r reviewRuleModelResolver) RowIndex() int32 {
	return int32(r.rule.RowIndex)
}

func (r reviewRuleModelResolver) ExactValue() *string {
	if r.rule.ExactValue == "" {
		return nil
	}
	return &r.rule.ExactValue
}

func (r reviewRuleModelResolver) GreaterThan() *string {
	if r.rule.GreaterThan == "" {
		return nil
	}
	return &r.rule.GreaterThan
}

func (r reviewRuleModelResolver) LessThan() *string {
	if r.rule.LessThan == "" {
		return nil
	}
	return &r.rule.LessThan
}

func (r reviewRuleModelResolver) ShouldBeNull() bool {
	return r.rule.ShouldBeNull
}
