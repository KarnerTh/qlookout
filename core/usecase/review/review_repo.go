package review

type ReviewRepo interface {
	GetRules(lookoutId int) ([]ReviewRule, error)
}
