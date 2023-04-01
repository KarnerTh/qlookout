package review

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KarnerTh/query-lookout/usecase/query"
	"github.com/KarnerTh/query-lookout/usecase/watch"
)

type testCaseData struct {
	Rows              []map[string]any
	rule              ReviewRule
	expectedResult    bool
	shouldReturnError bool
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
		ColumnType:  Int,
	},
	expectedResult: true,
}

var greaterThanFailure = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "2",
		ColumnType:  Int,
	},
	expectedResult: false,
}

var lessThanSuccess = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "2",
		ColumnType: Int,
	},
	expectedResult: true,
}

var lessThanFailure = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "1",
		ColumnType: Int,
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
		ColumnType:  Int,
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
		ColumnType:  Int,
	},
	expectedResult: false,
}

// no rules are provided
var undefinedRule = testCaseData{
	Rows: []map[string]any{{"column1": int64(1)}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
	},
	expectedResult:    false,
	shouldReturnError: true,
}

var exactValueFloatTypeSuccess = testCaseData{
	Rows: []map[string]any{{"column1": 1.123}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		ExactValue: "1.123",
	},
	expectedResult: true,
}

var exactValueFloatTypeFailure = testCaseData{
	Rows: []map[string]any{{"column1": 1.1234}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		ExactValue: "1.123",
	},
	expectedResult: false,
}

var greaterThanValueFloatTypeSuccess = testCaseData{
	Rows: []map[string]any{{"column1": 1.123}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "1.12",
		ColumnType:  Float,
	},
	expectedResult: true,
}

var greaterThanValueFloatTypeFailure = testCaseData{
	Rows: []map[string]any{{"column1": 1.12}},
	rule: ReviewRule{
		RowIndex:    0,
		ColumnName:  "column1",
		GreaterThan: "1.123",
		ColumnType:  Float,
	},
	expectedResult: false,
}

var lessThanValueFloatTypeSuccess = testCaseData{
	Rows: []map[string]any{{"column1": 1.12}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "1.123",
		ColumnType: Float,
	},
	expectedResult: true,
}

var lessThanValueFloatTypeFailure = testCaseData{
	Rows: []map[string]any{{"column1": 1.123}},
	rule: ReviewRule{
		RowIndex:   0,
		ColumnName: "column1",
		LessThan:   "1.12",
		ColumnType: Float,
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
		exactValueFloatTypeSuccess,
		exactValueFloatTypeFailure,
		greaterThanValueFloatTypeSuccess,
		greaterThanValueFloatTypeFailure,
		lessThanValueFloatTypeSuccess,
		lessThanValueFloatTypeFailure,
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
			result, err := validate(watchResult, rule)

			// Assert
			if testCase.shouldReturnError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
