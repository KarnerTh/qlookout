package review

import (
	"fmt"
	"strconv"

	"github.com/KarnerTh/query-lookout/usecase/watch"
	log "github.com/sirupsen/logrus"
)

func validate(watchResult watch.WatchResult, rule ReviewRule) bool {
	// TODO: use column type to parse correctly?
	actualValue := watchResult.Result.Rows[rule.RowIndex][rule.ColumnName]

	if rule.ExactValue != "" {
		value := fmt.Sprint(actualValue)
		expectedValue := rule.ExactValue
		return value == expectedValue
	}

	if rule.ShouldBeNull {
		return actualValue == nil
	}

	// TODO: parse int correctly (not always int64?)
	// HACK: types of column need to be checked
	if rule.GreaterThan != "" && rule.LessThan == "" {
		value := actualValue.(int64)
		greaterThan, err := strconv.ParseInt(rule.GreaterThan, 10, 64)
		if err != nil {
			log.WithError(err).Errorf("Could not parse value %s", rule.GreaterThan)
			return false
		}

		return value > greaterThan
	}

	if rule.LessThan != "" && rule.GreaterThan == "" {
		value := actualValue.(int64)
		lessThan, err := strconv.ParseInt(rule.LessThan, 10, 64)
		if err != nil {
			log.WithError(err).Errorf("Could not parse value %s", rule.LessThan)
			return false
		}

		return value < lessThan
	}

	if rule.LessThan != "" && rule.GreaterThan != "" {
		value := actualValue.(int64)
		lessThan, err := strconv.ParseInt(rule.LessThan, 10, 64)
		if err != nil {
			log.WithError(err).Errorf("Could not parse value %s", rule.LessThan)
			return false
		}

		greaterThan, err := strconv.ParseInt(rule.GreaterThan, 10, 64)
		if err != nil {
			log.WithError(err).Errorf("Could not parse value %s", rule.GreaterThan)
			return false
		}

		return (value < lessThan) && (value > greaterThan)
	}

	log.Warnf("Rule with id %d has no validation parameters", rule.Id)
	return false
}
