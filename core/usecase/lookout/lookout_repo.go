package lookout

type LookoutRepo interface {
	Get() ([]LookoutConfig, error)
	GetById(id int) (*LookoutConfig, error)
	Create(data LookoutConfigCreate) (*LookoutConfig, error)
	Update(id int, data LookoutConfigUpdate) (*LookoutConfig, error)
}
