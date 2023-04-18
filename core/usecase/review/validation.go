package review

import (
	"fmt"
	"strconv"

	"github.com/KarnerTh/query-lookout/usecase/watch"
)

func validate(watchResult watch.WatchResult, rule ReviewRule) (bool, error) {
	if len(watchResult.Result.Rows) < rule.RowIndex+1 {
		return false, fmt.Errorf("Row index is larger than row count (result length: %d, row index: %d)", len(watchResult.Result.Rows), rule.RowIndex)
	}

	actualValue, ok := watchResult.Result.Rows[rule.RowIndex][rule.ColumnName]
	if !ok {
		return false, fmt.Errorf("Rule column not found in result (%s)", rule.ColumnName)
	}

	if rule.ExactValue != "" {
		value := fmt.Sprint(actualValue)
		expectedValue := rule.ExactValue
		return value == expectedValue, nil
	}

	if rule.ShouldBeNull {
		return actualValue == nil, nil
	}

	rangeResult := true

	if rule.GreaterThan != "" {
		if rule.ColumnType == Int {
			value, rule, err := getInt64Values(actualValue, rule.GreaterThan)
			if err != nil {
				return false, err
			}
			rangeResult = rangeResult && value > rule
		} else if rule.ColumnType == Float {
			value, rule, err := getFloat64Values(actualValue, rule.GreaterThan)
			if err != nil {
				return false, err
			}
			rangeResult = rangeResult && value > rule
		} else {
			rangeResult = false
		}
	}

	if rule.LessThan != "" {
		if rule.ColumnType == Int {
			value, rule, err := getInt64Values(actualValue, rule.LessThan)
			if err != nil {
				return false, err
			}
			rangeResult = rangeResult && value < rule
		} else if rule.ColumnType == Float {
			value, rule, err := getFloat64Values(actualValue, rule.LessThan)
			if err != nil {
				return false, err
			}
			rangeResult = rangeResult && value < rule
		} else {
			rangeResult = false
		}
	}

	if rule.GreaterThan != "" || rule.LessThan != "" {
		return rangeResult, nil
	}

	return false, fmt.Errorf("Rule with id %d has no validation parameters", rule.Id)
}

func getInt64Values(actualValue any, ruleValue string) (actualValueResult int64, ruleValueResult int64, error error) {
	actualValueResult, ok := actualValue.(int64)
	if !ok {
		return 0, 0, fmt.Errorf("Could not parse value %v as int64", actualValue)
	}

	ruleValueResult, err := strconv.ParseInt(ruleValue, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Could not parse value %s", ruleValue)
	}

	return actualValueResult, ruleValueResult, nil
}

func getFloat64Values(actualValue any, ruleValue string) (actualValueResult float64, ruleValueResult float64, error error) {
	actualValueResult, ok := actualValue.(float64)
	if !ok {
		return 0, 0, fmt.Errorf("Could not parse value %v as float64", actualValue)
	}

	ruleValueResult, err := strconv.ParseFloat(ruleValue, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Could not parse value %s", ruleValue)
	}

	return actualValueResult, ruleValueResult, nil
}
