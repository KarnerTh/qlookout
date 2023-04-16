package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/KarnerTh/query-lookout/usecase/review"
)

type reviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) review.ReviewRepo {
	return reviewRepo{db: db}
}

func (r reviewRepo) GetForLookout(lookoutId int) ([]review.ReviewRule, error) {
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
		rule, err := scanRule(rows)
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (r reviewRepo) GetById(id int) (review.ReviewRule, error) {
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
    where id = ?
    `, id)
	if err != nil {
		return review.ReviewRule{}, err
	}

	if !rows.Next() {
		return review.ReviewRule{}, fmt.Errorf("No rule found for id %d", id)
	}

	return scanRule(rows)
}

func (r reviewRepo) Create(data review.ReviewRuleCreate) (review.ReviewRule, error) {
	var id int
	err := r.db.QueryRow(`
insert into review_rule(lookout_id, column_name, column_type, row_index, exact_value, greater_than, less_than,
                        should_be_null)
values (?, ?, ?, ?, ?, ?, ?, ?)
returning id
    `, data.LookoutId, data.ColumnName, data.ColumnType, data.RowIndex, data.ExactValue, data.GreaterThan, data.LessThan, data.ShouldBeNull).Scan(&id)

	if err != nil {
		return review.ReviewRule{}, err
	}

	return r.GetById(id)
}

func scanRule(row *sql.Rows) (review.ReviewRule, error) {
	var rule review.ReviewRule
	var exactValue sql.NullString
	var greaterThan sql.NullString
	var lessThan sql.NullString
	var shouldBeNull sql.NullBool

	if err := row.Scan(
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
		return review.ReviewRule{}, err
	}

	rule.ExactValue = exactValue.String
	rule.GreaterThan = greaterThan.String
	rule.LessThan = lessThan.String
	rule.ShouldBeNull = shouldBeNull.Bool

	return rule, nil
}
