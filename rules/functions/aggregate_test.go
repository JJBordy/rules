package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregateFunctions(t *testing.T) {

	af := AllAggregateFunctions()

	testCases := []struct {
		funcName string
		inputs   []any
		expected float64
	}{
		{
			funcName: MIN,
			inputs:   []any{0, 1, 2, 3},
			expected: 0,
		},
		{
			funcName: MAX,
			inputs:   []any{1, 2, 3, 4},
			expected: 4,
		},
		{
			funcName: SUM,
			inputs:   []any{1, 2, 3, 4},
			expected: 10,
		},
		{
			funcName: AVG,
			inputs:   []any{1, 2, 3, 4},
			expected: 2.5,
		},
		{
			funcName: COUNT,
			inputs:   []any{0, 1, 2, 3},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.funcName, func(t *testing.T) {
			af, ok := af.GetFunction(tc.funcName)
			assert.True(t, ok)

			result, err := af(tc.inputs)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
