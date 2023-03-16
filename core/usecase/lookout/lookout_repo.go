package lookout

type LookoutRepo interface {
	Get() ([]LookoutConfig, error)
}
