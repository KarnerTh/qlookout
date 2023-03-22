package lookout

type LookoutRepo interface {
	GetConfigs() ([]LookoutConfig, error)
	GetConfig(id int) (*LookoutConfig, error)
}
