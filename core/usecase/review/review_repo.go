package review

type ReviewRepo interface {
	GetForLookout(lookoutId int) ([]ReviewRule, error)
	GetById(id int) (*ReviewRule, error)
	Create(data ReviewRuleCreate) (*ReviewRule, error)
	Update(id int, data ReviewRuleUpdate) (*ReviewRule, error)
	Delete(id int) (*ReviewRule, error)
}
