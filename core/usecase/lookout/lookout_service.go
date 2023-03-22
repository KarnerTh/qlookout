package lookout

import log "github.com/sirupsen/logrus"
// TODO: Service needed?

type LookoutService interface {
	GetConfigs() ([]LookoutConfig, error)
	GetConfig(id int) (*LookoutConfig, error)
}

type lookoutService struct {
	lookoutRepo LookoutRepo
}

func NewLookoutService(lookoutRepo LookoutRepo) LookoutService {
	return &lookoutService{
		lookoutRepo: lookoutRepo,
	}
}

func (l lookoutService) GetConfigs() ([]LookoutConfig, error) {
	configs, err := l.lookoutRepo.GetConfigs()
	if err != nil {
		log.WithError(err).Error("Could not get lookout configs")
		return nil, err
	}

	return configs, nil
}

func (l lookoutService) GetConfig(id int) (*LookoutConfig, error) {
	config, err := l.lookoutRepo.GetConfig(id)
	if err != nil {
		log.WithError(err).Errorf("Could not get lookout config with id %d", id)
		return nil, err
	}

	return config, nil
}
