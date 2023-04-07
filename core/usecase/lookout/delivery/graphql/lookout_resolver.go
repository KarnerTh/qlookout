package graphql

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
)

type LookoutResolver struct {
	lookoutService lookout.LookoutService
}

func NewLookoutResolver(lookoutService lookout.LookoutService) LookoutResolver {
	return LookoutResolver{
		lookoutService: lookoutService,
	}
}

func (r LookoutResolver) Lookouts() []lookoutConfigModel {
	lookouts, err := r.lookoutService.GetConfigs()
	if err != nil {
		// TODO: send error
	}

	return lookoutToModel(lookouts)
}

func lookoutToModel(domain []lookout.LookoutConfig) []lookoutConfigModel {
	models := make([]lookoutConfigModel, len(domain))
	for i, value := range domain {
		models[i] = lookoutConfigModel{
			Id:          int32(value.Id),
			Name:        value.Name,
			Query:       value.Query,
			Cron:        value.Cron,
			NotifyLocal: value.NotifyLocal,
			NotifyMail:  value.NotifyMail,
		}
	}

	return models
}
