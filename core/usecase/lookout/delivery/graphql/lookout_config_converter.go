package graphql

import "github.com/KarnerTh/query-lookout/usecase/lookout"

func configToModel(domain lookout.LookoutConfig) lookoutConfigModel {
	return lookoutConfigModel{
		Id:          int32(domain.Id),
		Name:        domain.Name,
		Query:       domain.Query,
		Cron:        domain.Cron,
		NotifyLocal: domain.NotifyLocal,
		NotifyMail:  domain.NotifyMail,
	}
}
