package infrastructure

import (
	"database/sql"

	"github.com/KarnerTh/query-lookout/usecase/review"
)

type reviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) review.ReviewRepo {
	return &reviewRepo{db: db}
}

func (r reviewRepo) GetRules(lookoutId int) ([]review.ReviewRule, error) {
	rows, err := r.db.Query("select id, column_name, row_index, exact_value from review_rule")
	if err != nil {
		return nil, err
	}

	var rules []review.ReviewRule
	for rows.Next() {
		var rule review.ReviewRule
		if err = rows.Scan(&rule.Id, &rule.ColumnName, &rule.RowIndex, &rule.ExactValue); err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
