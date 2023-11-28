package graphql

import (
	"github.com/KarnerTh/qlookout/core/usecase/lookout"
	reviewGraphQl "github.com/KarnerTh/qlookout/core/usecase/review/delivery/graphql"
)

type LookoutResolver struct {
	lookoutManager lookout.LookoutManager
	lookoutRepo    lookout.LookoutRepo
	reviewResolver reviewGraphQl.ReviewResolver
}

func NewLookoutResolver(lookoutManager lookout.LookoutManager, lookoutRepo lookout.LookoutRepo, reviewResolver reviewGraphQl.ReviewResolver) LookoutResolver {
	return LookoutResolver{
		lookoutManager: lookoutManager,
		lookoutRepo:    lookoutRepo,
		reviewResolver: reviewResolver,
	}
}

func (r LookoutResolver) Lookouts() ([]lookoutConfigModel, error) {
	lookouts, err := r.lookoutRepo.Get()
	if err != nil {
		return nil, err
	}

	models := make([]lookoutConfigModel, len(lookouts))
	for i, value := range lookouts {
		models[i] = lookoutConfigModelResolver{lookout: value, reviewResolver: r.reviewResolver}
	}

	return models, nil
}

func (r LookoutResolver) CreateLookout(args struct{ Data lookoutConfigCreateModel }) (lookoutConfigModel, error) {
	data, err := r.lookoutRepo.Create(lookout.LookoutConfigCreate{
		Name:        args.Data.Name,
		Cron:        args.Data.Cron,
		Query:       args.Data.Query,
		NotifyLocal: args.Data.NotifyLocal,
		NotifyMail:  args.Data.NotifyMail,
	})

	if err != nil {
		return nil, err
	}

	err = r.lookoutManager.Watch(data.Id)
	if err != nil {
		return nil, err
	}

	return lookoutConfigModelResolver{lookout: *data, reviewResolver: r.reviewResolver}, nil
}

func (r LookoutResolver) Lookout(args struct{ Id int32 }) (lookoutConfigModel, error) {
	data, err := r.lookoutRepo.GetById(int(args.Id))
	if err != nil {
		return nil, err
	}

	return lookoutConfigModelResolver{lookout: *data, reviewResolver: r.reviewResolver}, nil
}

func (r LookoutResolver) UpdateLookout(args struct {
	Id   int32
	Data lookoutConfigUpdateModel
}) (lookoutConfigModel, error) {
	data, err := r.lookoutRepo.Update(
		int(args.Id),
		lookout.LookoutConfigUpdate{
			Name:        args.Data.Name,
			Cron:        args.Data.Cron,
			Query:       args.Data.Query,
			NotifyLocal: args.Data.NotifyLocal,
			NotifyMail:  args.Data.NotifyMail,
		})

	if err != nil {
		return nil, err
	}

	err = r.lookoutManager.Reload(int(args.Id))
	if err != nil {
		return nil, err
	}

	return lookoutConfigModelResolver{lookout: *data, reviewResolver: r.reviewResolver}, nil
}

func (r LookoutResolver) DeleteLookout(args struct{ Id int32 }) (lookoutConfigModel, error) {
	data, err := r.lookoutRepo.Delete(int(args.Id))
	if err != nil {
		return nil, err
	}

	err = r.lookoutManager.Remove(int(args.Id))
	if err != nil {
		return nil, err
	}

	return lookoutConfigModelResolver{lookout: *data, reviewResolver: r.reviewResolver}, nil
}
