package review

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/KarnerTh/query-lookout/core/usecase/watch"
)

type ValidationResult struct {
	IsValid     bool
	Description string
}

func validate(watchResult watch.WatchResult, rule ReviewRule) (ValidationResult, error) {
	if len(watchResult.Result.Rows) < rule.RowIndex+1 {
		return ValidationResult{IsValid: false}, fmt.Errorf("Row index is larger than row count (result length: %d, row index: %d)", len(watchResult.Result.Rows), rule.RowIndex)
	}

	actualValue, ok := watchResult.Result.Rows[rule.RowIndex][rule.ColumnName]
	if !ok {
		return ValidationResult{IsValid: false}, fmt.Errorf("Rule column not found in result (%s)", rule.ColumnName)
	}

	if rule.ExactValue != "" {
		value := fmt.Sprint(actualValue)
		expectedValue := rule.ExactValue

		isValid := value == expectedValue
		if isValid {
			return ValidationResult{IsValid: true}, nil
		}

		return ValidationResult{IsValid: false, Description: fmt.Sprintf("Excpected exact value '%s', but got '%s'", expectedValue, value)}, nil
	}

	if rule.ShouldBeNull {
		isValid := actualValue == nil
		if isValid {
			return ValidationResult{IsValid: true}, nil
		}

		return ValidationResult{IsValid: false, Description: fmt.Sprintf("Expected value to be null, but got '%s'", fmt.Sprint(actualValue))}, nil
	}

	rangeResult := true
	rangeDescriptions := []string{}
	var greaterRangeError error
	var lessRangeError error

	if rule.GreaterThan != "" {
		var greaterIsValid bool
		if rule.ColumnType == Int {
			value, rule, err := getInt64Values(actualValue, rule.GreaterThan)
			if err != nil {
				return ValidationResult{IsValid: false}, err
			}
			greaterIsValid = value > rule
		} else if rule.ColumnType == Float {
			value, rule, err := getFloat64Values(actualValue, rule.GreaterThan)
			if err != nil {
				return ValidationResult{IsValid: false}, err
			}

			greaterIsValid = value > rule
		} else {
			greaterRangeError = fmt.Errorf("Greater than not supported with this column type: %s", rule.ColumnType)
			greaterIsValid = false
		}

		rangeResult = rangeResult && greaterIsValid
		if !greaterIsValid {
			rangeDescriptions = append(rangeDescriptions, fmt.Sprintf("Expected value to be greater than '%s', but got '%s'", rule.GreaterThan, fmt.Sprint(actualValue)))
		}
	}

	if rule.LessThan != "" {
		var lessThanIsValid bool
		if rule.ColumnType == Int {
			value, rule, err := getInt64Values(actualValue, rule.LessThan)
			if err != nil {
				return ValidationResult{IsValid: false}, err
			}
			lessThanIsValid = value < rule
		} else if rule.ColumnType == Float {
			value, rule, err := getFloat64Values(actualValue, rule.LessThan)
			if err != nil {
				return ValidationResult{IsValid: false}, err
			}
			lessThanIsValid = value < rule
		} else {
			lessRangeError = fmt.Errorf("Less than not supported with this column type: %s", rule.ColumnType)
			lessThanIsValid = false
		}

		rangeResult = rangeResult && lessThanIsValid
		if !lessThanIsValid {
			rangeDescriptions = append(rangeDescriptions, fmt.Sprintf("Expected value to be less than '%s', but got '%s'", rule.LessThan, fmt.Sprint(actualValue)))
		}
	}

	if rule.GreaterThan != "" || rule.LessThan != "" {
		return ValidationResult{IsValid: rangeResult, Description: strings.Join(rangeDescriptions, ",")}, errors.Join(greaterRangeError, lessRangeError)
	}

	return ValidationResult{IsValid: false}, fmt.Errorf("Rule with id %d has no validation parameters", rule.Id)
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
