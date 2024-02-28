package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultSingleInputFunctions(t *testing.T) {

	defaultFunctions := Default()

	testCases := []struct {
		FuncName     string
		Args         []any
		Input        any
		ExpectOutput bool
		ExpectError  bool
	}{
		{
			FuncName:     Empty,
			Args:         []any{},
			Input:        "",
			ExpectOutput: true,
		},
		{
			FuncName:     Empty,
			Args:         []any{},
			Input:        "SOMETHING",
			ExpectOutput: false,
		},

		{
			FuncName:     NonEmpty,
			Args:         []any{},
			Input:        "",
			ExpectOutput: false,
		},
		{
			FuncName:     NonEmpty,
			Args:         []any{},
			Input:        "SOMETHING",
			ExpectOutput: true,
		},

		{
			FuncName:     Equal,
			Args:         []any{"SOMETHING"},
			Input:        "SOMETHING",
			ExpectOutput: true,
		},
		{
			FuncName:     Equal,
			Args:         []any{"SOMETHING"},
			Input:        "SOMETHINGELSE",
			ExpectOutput: false,
		},
		{
			FuncName:     Greater,
			Args:         []any{1},
			Input:        2,
			ExpectOutput: true,
		},
		{
			FuncName:     Greater,
			Args:         []any{1},
			Input:        0,
			ExpectOutput: false,
		},
		{
			FuncName:     Greater,
			Args:         []any{1},
			Input:        1,
			ExpectOutput: false,
		},
		{
			FuncName:     GreaterEq,
			Args:         []any{1},
			Input:        2,
			ExpectOutput: true,
		},
		{
			FuncName:     GreaterEq,
			Args:         []any{1},
			Input:        1,
			ExpectOutput: true,
		},
		{
			FuncName:     GreaterEq,
			Args:         []any{1},
			Input:        0,
			ExpectOutput: false,
		},

		{
			FuncName:     Lower,
			Args:         []any{1},
			Input:        0,
			ExpectOutput: true,
		},
		{
			FuncName:     Lower,
			Args:         []any{1},
			Input:        1,
			ExpectOutput: false,
		},
		{
			FuncName:     Lower,
			Args:         []any{1},
			Input:        2,
			ExpectOutput: false,
		},
		{
			FuncName:     LowerEq,
			Args:         []any{1},
			Input:        0,
			ExpectOutput: true,
		},
		{
			FuncName:     LowerEq,
			Args:         []any{1},
			Input:        1,
			ExpectOutput: true,
		},
		{
			FuncName:     LowerEq,
			Args:         []any{1},
			Input:        2,
			ExpectOutput: false,
		},
		{
			FuncName:     Between,
			Args:         []any{1, 2},
			Input:        1,
			ExpectOutput: false,
		},
		{
			FuncName:     Between,
			Args:         []any{1, 3},
			Input:        2,
			ExpectOutput: true,
		},
		{
			FuncName:     Between,
			Args:         []any{1, 2},
			Input:        3,
			ExpectOutput: false,
		},
		{
			FuncName:     BetweenEq,
			Args:         []any{1, 2},
			Input:        1,
			ExpectOutput: true,
		},
		{
			FuncName:     BetweenEq,
			Args:         []any{1, 3},
			Input:        2,
			ExpectOutput: true,
		},
		{
			FuncName:     BetweenEq,
			Args:         []any{1, 3},
			Input:        3,
			ExpectOutput: true,
		},
		{
			FuncName:     BetweenEq,
			Args:         []any{1, 3},
			Input:        4,
			ExpectOutput: false,
		},
		{
			FuncName:     BetweenEq,
			Args:         []any{1, 3},
			Input:        0,
			ExpectOutput: false,
		},
		{
			FuncName:     NotBetween,
			Args:         []any{1, 2},
			Input:        3,
			ExpectOutput: true,
		},
		{
			FuncName:     NotBetween,
			Args:         []any{1, 2},
			Input:        2,
			ExpectOutput: false,
		},
		{
			FuncName:     NotBetween,
			Args:         []any{1, 2},
			Input:        1,
			ExpectOutput: false,
		},
		{
			FuncName:     NotBetweenEq,
			Args:         []any{1, 2},
			Input:        3,
			ExpectOutput: true,
		},
		{
			FuncName:     NotBetweenEq,
			Args:         []any{1, 2},
			Input:        2,
			ExpectOutput: false,
		},
		{
			FuncName:     NotBetweenEq,
			Args:         []any{1, 2},
			Input:        1,
			ExpectOutput: false,
		},
		{
			FuncName:     EqualIgnoreCase,
			Args:         []any{"a", "a"},
			Input:        "A",
			ExpectOutput: true,
		},
		{
			FuncName:     EqualIgnoreCase,
			Args:         []any{"a", "a"},
			Input:        "b",
			ExpectOutput: false,
		},
		{
			FuncName:     EqualAny,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "ABC",
			ExpectOutput: true,
		},
		{
			FuncName:     EqualAny,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "DEF",
			ExpectOutput: false,
		},
	}

	/*
		// Greater - the input is a number and is greater than the argument
		Greater = "Greater"
		// GreaterEq - the input is a number and is greater or equal to the argument
		GreaterEq = "GreaterEq"
		// Lower - the input is a number and is lower than the argument
		Lower = "Lower"
		// LowerEq - the input is a number and is lower or equal to the argument
		LowerEq = "LowerEq"
		// Between - the input is a number and is between the two arguments
		Between = "Between"
		// BetweenEq - the input is a number and is between or equal to the two arguments
		BetweenEq = "BetweenEq"
		// NotBetween - the input is a number and is not between the two arguments
		NotBetween = "NotBetween"
		// NotBetweenEq - the input is a number and is not between or equal to the two arguments
		NotBetweenEq = "NotBetweenEq"

		// EqualIgnoreCase - the input is equal to the argument; case-insensitive
		EqualIgnoreCase = "EqualIgnoreCase"
		// EqualAny - the input is equal to any of the arguments
		EqualAny = "EqualAny"
		// NotEqualAny - the input is not equal to any of the arguments
		NotEqualAny = "NotEqualAny"
		// StartsWith - the input starts with the argument
		StartsWith = "StartsWith"
		// StartsWithIgnoreCase - the input starts with the argument; case-insensitive
		StartsWithIgnoreCase = "StartsWithIgnoreCase"
		// EndsWith - the input ends with the argument
		EndsWith = "EndsWith"
		// EndsWithIgnoreCase - the input ends with the argument; case-insensitive
		EndsWithIgnoreCase = "EndsWithIgnoreCase"
		// Contains - the input contains the argument
		Contains = "Contains"
		// ContainsIgnoreCase - the input contains the argument; case-insensitive
		ContainsIgnoreCase = "ContainsIgnoreCase"
	*/

	for _, tc := range testCases {
		t.Run(tc.FuncName, func(t *testing.T) {
			result, err := defaultFunctions[tc.FuncName](tc.Input, tc.Args)
			assert.Equal(t, tc.ExpectOutput, result, fmt.Sprintf("%+v\n", tc))
			if tc.ExpectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err, fmt.Sprintf("%+v\n", tc), fmt.Sprint(err))
			}
		})
	}
}
