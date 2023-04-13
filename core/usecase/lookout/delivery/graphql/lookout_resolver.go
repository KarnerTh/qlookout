package graphql

import (
	"github.com/KarnerTh/query-lookout/usecase/lookout"
)

type LookoutResolver struct {
	lookoutManager lookout.LookoutManager
	lookoutService lookout.LookoutService
}

func NewLookoutResolver(lookoutManager lookout.LookoutManager, lookoutService lookout.LookoutService) LookoutResolver {
	return LookoutResolver{
		lookoutManager: lookoutManager,
		lookoutService: lookoutService,
	}
}

func (r LookoutResolver) Lookouts() ([]lookoutConfigModel, error) {
	lookouts, err := r.lookoutService.Get()
	if err != nil {
		return nil, err
	}

	return lookoutToModel(lookouts), nil
}

func lookoutToModel(domain []lookout.LookoutConfig) []lookoutConfigModel {
	models := make([]lookoutConfigModel, len(domain))
	for i, value := range domain {
		models[i] = configToModel(value)
	}

	return models
}

func (r LookoutResolver) CreateLookout(args struct{ Data lookoutConfigCreateModel }) (lookoutConfigModel, error) {
	data, err := r.lookoutService.Create(lookout.LookoutConfigCreate{
		Name:        args.Data.Name,
		Cron:        args.Data.Cron,
		Query:       args.Data.Query,
		NotifyLocal: args.Data.NotifyLocal,
		NotifyMail:  args.Data.NotifyMail,
	})

	if err != nil {
		return lookoutConfigModel{}, err
	}

	r.lookoutManager.Watch(data.Id)
	return configToModel(*data), nil
}

func (r LookoutResolver) Lookout(args struct{ Id int32 }) (lookoutConfigModel, error) {
	data, err := r.lookoutService.GetById(int(args.Id))
	if err != nil {
		return lookoutConfigModel{}, err
	}

	return configToModel(*data), nil
}
