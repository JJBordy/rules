package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllListFunctionConstraints(t *testing.T) {
	allListFunctionConstraints := AllListFunctionConstraints()

	testCases := []struct {
		constraintName string
		listTotal      int
		passedTotal    int
		args           []int
		expectedResult bool
	}{
		{
			constraintName: All,
			listTotal:      5,
			passedTotal:    5,
			expectedResult: true,
		},
		{
			constraintName: All,
			listTotal:      5,
			passedTotal:    4,
			expectedResult: false,
		},
		{
			constraintName: All,
			listTotal:      0,
			passedTotal:    0,
			expectedResult: true,
		},
		{
			constraintName: AtLeast,
			passedTotal:    3,
			args:           []int{3},
			expectedResult: true,
		},
		{
			constraintName: AtLeast,
			passedTotal:    2,
			args:           []int{3},
			expectedResult: false,
		},
		{
			constraintName: AtLeast,
			passedTotal:    0,
			args:           []int{1},
			expectedResult: false,
		},
		{
			constraintName: AtMost,
			passedTotal:    3,
			args:           []int{3},
			expectedResult: true,
		},
		{
			constraintName: AtMost,
			passedTotal:    2,
			args:           []int{3},
			expectedResult: true,
		},
		{
			constraintName: AtMost,
			passedTotal:    4,
			args:           []int{3},
			expectedResult: false,
		},
		{
			constraintName: AtMost,
			passedTotal:    4,
			args:           []int{0},
			expectedResult: false,
		},
		{
			constraintName: Exactly,
			passedTotal:    3,
			args:           []int{3},
			expectedResult: true,
		},
		{
			constraintName: Exactly,
			passedTotal:    0,
			args:           []int{0},
			expectedResult: true,
		},
		{
			constraintName: Exactly,
			passedTotal:    3,
			args:           []int{4},
			expectedResult: false,
		},
		{
			constraintName: None,
			passedTotal:    10,
			expectedResult: false,
		},
		{
			constraintName: None,
			passedTotal:    0,
			expectedResult: true,
		},
		{
			constraintName: AtLeastFraction,
			listTotal:      8,
			passedTotal:    4,
			args:           []int{2, 4},
			expectedResult: true,
		},
		{
			constraintName: AtLeastFraction,
			listTotal:      8,
			passedTotal:    3,
			args:           []int{2, 4},
			expectedResult: false,
		},
		{
			constraintName: AtLeastFraction,
			listTotal:      8,
			passedTotal:    7,
			args:           []int{3, 4},
			expectedResult: true,
		},
		{
			constraintName: AtLeastFraction,
			listTotal:      8,
			passedTotal:    4,
			args:           []int{3, 4},
			expectedResult: false,
		},
		{
			constraintName: AtMostFraction,
			listTotal:      8,
			passedTotal:    4,
			args:           []int{2, 4},
			expectedResult: true,
		},
		{
			constraintName: AtMostFraction,
			listTotal:      8,
			passedTotal:    3,
			args:           []int{2, 4},
			expectedResult: true,
		},
		{
			constraintName: AtMostFraction,
			listTotal:      8,
			passedTotal:    7,
			args:           []int{3, 4},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.constraintName, func(t *testing.T) {
			gotResult := allListFunctionConstraints[tc.constraintName](tc.listTotal, tc.passedTotal, tc.args)
			assert.Equal(t, tc.expectedResult, gotResult)
		})
	}
}
