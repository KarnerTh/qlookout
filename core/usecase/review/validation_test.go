package review

import (
	"fmt"
	"testing"

	"github.com/KarnerTh/query-lookout/usecase/query"
	"github.com/KarnerTh/query-lookout/usecase/watch"
	"github.com/stretchr/testify/assert"
)

type testCaseData struct {
	Rows           []map[string]any
	rule           ReviewRule
	expectedResult bool
}

var exactValueSuccess = testCaseData{
	Rows: []map[string]any{{"column1": "exactValue"}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		ExactValue: "exactValue",
	},
	expectedResult: true,
}

var exactValueFailure = testCaseData{
	Rows: []map[string]any{{"column1": "exactValue"}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		ExactValue: "shouldFail",
	},
	expectedResult: false,
}

var shouldBeNullSuccess = testCaseData{
	Rows: []map[string]any{{"column1": nil}},
	rule: ReviewRule{
		RowIndex:     0,
		ColumnName:   "column1",
		ShouldBeNull: true,
	},
	expectedResult: true,
}

var shouldBeNullFailure = testCaseData{
	Rows: []map[string]any{{"column1": "notNull"}},
	rule: ReviewRule{
		RowIndex:     0,
		ColumnName:   "column1",
		ShouldBeNull: true,
	},
	expectedResult: false,
}

var greaterThanSuccess = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "0",
	},
	expectedResult: true,
}

var greaterThanFailure = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "2",
	},
	expectedResult: false,
}

var lessThanSuccess = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "2",
	},
	expectedResult: true,
}

var lessThanFailure = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "1",
	},
	expectedResult: false,
}

var greaterAndLessThanSuccess = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "0",
		LessThan:    "2",
	},
	expectedResult: true,
}

var greaterAndLessThanFailure = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "2",
		LessThan:    "1",
	},
	expectedResult: false,
}

var undefinedRule = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
	},
	expectedResult: false,
}

func TestValidate(t *testing.T) {
	testCases := []testCaseData{
		exactValueSuccess,
		exactValueFailure,
		shouldBeNullSuccess,
		shouldBeNullFailure,
		greaterThanSuccess,
		greaterThanFailure,
		lessThanSuccess,
		lessThanFailure,
		greaterAndLessThanSuccess,
		greaterAndLessThanFailure,
    undefinedRule,
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run #%d", i), func(t *testing.T) {
			// Arrange
			watchResult := watch.WatchResult{
				Result: query.QueryResult{
					Rows: testCase.Rows,
				},
			}
			rule := testCase.rule

			// Act
			result := validate(watchResult, rule)

			// Assert
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
