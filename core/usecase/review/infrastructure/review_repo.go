package infrastructure

import (
	"database/sql"
	"fmt"
	"strings"

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

		rules = append(rules, *rule)
	}

	return rules, nil
}

func (r reviewRepo) GetById(id int) (*review.ReviewRule, error) {
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

	// if rows.Next is not called til it returns false, the rows are not automatically closed
	// source: https://pkg.go.dev/database/sql#Rows.Close
	defer func() {
		_ = rows.Close()
	}()

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("No rule found for id %d", id)
	}

	return scanRule(rows)
}

func (r reviewRepo) Create(data review.ReviewRuleCreate) (*review.ReviewRule, error) {
	var id int
	err := r.db.QueryRow(`
insert into review_rule(lookout_id, column_name, column_type, row_index, exact_value, greater_than, less_than,
                        should_be_null)
values (?, ?, ?, ?, ?, ?, ?, ?)
returning id
    `, data.LookoutId, data.ColumnName, data.ColumnType, data.RowIndex, data.ExactValue, data.GreaterThan, data.LessThan, data.ShouldBeNull).Scan(&id)

	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func (r reviewRepo) Update(id int, data review.ReviewRuleUpdate) (*review.ReviewRule, error) {
	var updateProps []string
	var args []any

	if data.ColumnName != nil {
		updateProps = append(updateProps, "column_name=?")
		args = append(args, data.ColumnName)
	}
	if data.ColumnType != nil {
		updateProps = append(updateProps, "column_type=?")
		args = append(args, data.ColumnType)
	}
	if data.RowIndex != nil {
		updateProps = append(updateProps, "row_index=?")
		args = append(args, data.RowIndex)
	}
	if data.ExactValue != nil {
		updateProps = append(updateProps, "exact_value=?")
		args = append(args, data.ExactValue)
	}
	if data.GreaterThan != nil {
		updateProps = append(updateProps, "greater_than=?")
		args = append(args, data.GreaterThan)
	}
	if data.LessThan != nil {
		updateProps = append(updateProps, "less_than=?")
		args = append(args, data.LessThan)
	}
	if data.ShouldBeNull != nil {
		updateProps = append(updateProps, "should_be_null=?")
		args = append(args, data.ShouldBeNull)
	}

	if len(updateProps) == 0 {
		// nothing to Update
		return r.GetById(id)
	}

	query := "update review_rule set " + strings.Join(updateProps, ",") + "where id = ?"
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}

func scanRule(row *sql.Rows) (*review.ReviewRule, error) {
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
		return nil, err
	}

	rule.ExactValue = exactValue.String
	rule.GreaterThan = greaterThan.String
	rule.LessThan = lessThan.String
	rule.ShouldBeNull = shouldBeNull.Bool

	return &rule, nil
}
