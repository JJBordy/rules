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
			Input:        "SOMETHING ELSE",
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
			ExpectOutput: true,
		},
		{
			FuncName:     NotBetween,
			Args:         []any{1, 2},
			Input:        1,
			ExpectOutput: true,
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
			Args:         []any{"aBc"},
			Input:        "AbC",
			ExpectOutput: true,
		},
		{
			FuncName:     EqualIgnoreCase,
			Args:         []any{"ABC"},
			Input:        "abcd",
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
		{
			FuncName:     EqualAnyIgnoreCase,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "AbC",
			ExpectOutput: true,
		},
		{
			FuncName:     EqualAnyIgnoreCase,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "def",
			ExpectOutput: false,
		},
		{
			FuncName:     NotEqualAny,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "DEF",
			ExpectOutput: true,
		},
		{
			FuncName:     NotEqualAny,
			Args:         []any{"a", "xyz", "ABC"},
			Input:        "ABC",
			ExpectOutput: false,
		},
		{
			FuncName:     StartsWith,
			Args:         []any{"Abr"},
			Input:        "Abracadabra",
			ExpectOutput: true,
		},
		{
			FuncName:     StartsWith,
			Args:         []any{"ABR"},
			Input:        "Abracadabra",
			ExpectOutput: false,
		},
		{
			FuncName:     StartsWithIgnoreCase,
			Args:         []any{"abr"},
			Input:        "ABRACADABRA",
			ExpectOutput: true,
		},
		{
			FuncName:     StartsWithIgnoreCase,
			Args:         []any{"abba"},
			Input:        "ABRACADABRA",
			ExpectOutput: false,
		},
		{
			FuncName:     EndsWith,
			Args:         []any{"bra"},
			Input:        "Abracadabra",
			ExpectOutput: true,
		},
		{
			FuncName:     EndsWith,
			Args:         []any{"BRA"},
			Input:        "Abracadabra",
			ExpectOutput: false,
		},
		{
			FuncName:     EndsWithIgnoreCase,
			Args:         []any{"bra"},
			Input:        "ABRACADABRA",
			ExpectOutput: true,
		},
		{
			FuncName:     EndsWithIgnoreCase,
			Args:         []any{"DARA"},
			Input:        "ABRACADABRA",
			ExpectOutput: false,
		},
		{
			FuncName:     Contains,
			Args:         []any{"cad"},
			Input:        "Abracadabra",
			ExpectOutput: true,
		},
		{
			FuncName:     Contains,
			Args:         []any{"CAD"},
			Input:        "Abracadabra",
			ExpectOutput: false,
		},
		{
			FuncName:     ContainsIgnoreCase,
			Args:         []any{"cad"},
			Input:        "ABRACADABRA",
			ExpectOutput: true,
		},
		{
			FuncName:     ContainsIgnoreCase,
			Args:         []any{"BAD"},
			Input:        "ABRACADABRA",
			ExpectOutput: false,
		},
	}

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
