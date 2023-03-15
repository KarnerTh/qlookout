package lookout

type LookoutRepo interface {
	Get() ([]Lookout, error)
}
