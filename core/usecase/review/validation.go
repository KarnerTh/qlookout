package review

import (
	"fmt"

	"github.com/KarnerTh/query-lookout/usecase/watch"
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

	return false
}
