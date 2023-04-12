package lookout

import log "github.com/sirupsen/logrus"

type LookoutService interface {
	Get() ([]LookoutConfig, error)
	GetById(id int) (*LookoutConfig, error)
	Create(data LookoutConfigCreate) (*LookoutConfig, error)
}

type lookoutService struct {
	lookoutRepo LookoutRepo
}

func NewLookoutService(lookoutRepo LookoutRepo) LookoutService {
	return lookoutService{
		lookoutRepo: lookoutRepo,
	}
}

func (l lookoutService) Get() ([]LookoutConfig, error) {
	configs, err := l.lookoutRepo.Get()
	if err != nil {
		log.WithError(err).Error("Could not get lookout configs")
		return nil, err
	}

	return configs, nil
}

func (l lookoutService) GetById(id int) (*LookoutConfig, error) {
	config, err := l.lookoutRepo.GetById(id)
	if err != nil {
		log.WithError(err).Errorf("Could not get lookout config with id %d", id)
		return nil, err
	}

	return config, nil
}

func (l lookoutService) Create(data LookoutConfigCreate) (*LookoutConfig, error) {
	result, err := l.lookoutRepo.Create(data)
	if err != nil {
		log.WithError(err).Error("Could not create lookout config ")
		return nil, err
	}

	return result, nil
}
