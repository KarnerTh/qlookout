package graphql

import (
	"github.com/KarnerTh/query-lookout/core/usecase/lookout"
	reviewGraphQl "github.com/KarnerTh/query-lookout/core/usecase/review/delivery/graphql"
)

type lookoutConfigModel interface {
	Id() int32
	Name() string
	Query() string
	Cron() string
	NotifyLocal() bool
	NotifyMail() bool
	Rules() ([]reviewGraphQl.ReviewRuleModel, error)
}

type lookoutConfigModelResolver struct {
	lookout        lookout.LookoutConfig
	reviewResolver reviewGraphQl.ReviewResolver
}

func (r lookoutConfigModelResolver) Id() int32 {
	return int32(r.lookout.Id)
}

func (r lookoutConfigModelResolver) Name() string {
	return r.lookout.Name
}

func (r lookoutConfigModelResolver) Query() string {
	return r.lookout.Query
}

func (r lookoutConfigModelResolver) Cron() string {
	return r.lookout.Cron
}

func (r lookoutConfigModelResolver) NotifyLocal() bool {
	return r.lookout.NotifyLocal
}

func (r lookoutConfigModelResolver) NotifyMail() bool {
	return r.lookout.NotifyMail
}

func (r lookoutConfigModelResolver) Rules() ([]reviewGraphQl.ReviewRuleModel, error) {
	return r.reviewResolver.Rules(struct{ LookoutId int32 }{LookoutId: int32(r.lookout.Id)})
}
