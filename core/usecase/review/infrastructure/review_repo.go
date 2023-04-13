package infrastructure

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/usecase/review"
)

type reviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) review.ReviewRepo {
	return reviewRepo{db: db}
}

func (r reviewRepo) GetRules(lookoutId int) ([]review.ReviewRule, error) {
	rows, err := r.db.Query(`
    select  id,
            lookout_id,
            column_name,
            column_type,
            row_index,
            exact_value,
            greater_than,
            less_than,
            should_be_null 
    from review_rule
    where lookout_id = ?
    `, lookoutId)
	if err != nil {
		return nil, err
	}

	var rules []review.ReviewRule
	for rows.Next() {
		var rule review.ReviewRule
		var exactValue sql.NullString
		var greaterThan sql.NullString
		var lessThan sql.NullString
		var shouldBeNull sql.NullBool

		if err := rows.Scan(
			&rule.Id,
			&rule.LookoutId,
			&rule.ColumnName,
			&rule.ColumnType,
			&rule.RowIndex,
			&exactValue,
			&greaterThan,
			&lessThan,
			&shouldBeNull,
		); err != nil {
			return nil, err
		}

		rule.ExactValue = exactValue.String
		rule.GreaterThan = greaterThan.String
		rule.LessThan = lessThan.String
		rule.ShouldBeNull = shouldBeNull.Bool

		rules = append(rules, rule)
	}

	return rules, nil
}
