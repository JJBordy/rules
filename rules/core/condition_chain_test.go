package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditionChains(t *testing.T) {

	t.Run("throws error on invalid chain type", func(t *testing.T) {
		_, err := NewConditionChain("invalid")
		assert.NotNil(t, err)
	})

	testCases := []struct {
		cType      string
		conditions []Condition
		expected   bool
	}{
		{
			cType: AND,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
			},
			expected: true,
		},
		{
			cType: AND,
			conditions: []Condition{
				conditionTrue{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: OR,
			conditions: []Condition{
				conditionFalse{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: OR,
			conditions: []Condition{
				conditionFalse{},
				conditionTrue{},
			},
			expected: true,
		},
		{
			cType: NAND,
			conditions: []Condition{
				conditionFalse{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: NAND,
			conditions: []Condition{
				conditionTrue{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: NAND,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
			},
			expected: false,
		},
		{
			cType: NOR,
			conditions: []Condition{
				conditionFalse{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: NOR,
			conditions: []Condition{
				conditionTrue{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: NOR,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
			},
			expected: false,
		},
		{
			cType: XOR,
			conditions: []Condition{
				conditionFalse{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: XOR,
			conditions: []Condition{
				conditionFalse{},
				conditionTrue{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: XOR,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: XOR,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
			},
			expected: false,
		},
		{
			cType: XNOR,
			conditions: []Condition{
				conditionFalse{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: XNOR,
			conditions: []Condition{
				conditionFalse{},
				conditionTrue{},
				conditionFalse{},
			},
			expected: false,
		},
		{
			cType: XNOR,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
				conditionFalse{},
			},
			expected: true,
		},
		{
			cType: XNOR,
			conditions: []Condition{
				conditionTrue{},
				conditionTrue{},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.cType, func(t *testing.T) {
			cc, _ := NewConditionChain(tc.cType)
			result, _ := cc.EvaluateConditions(map[string]interface{}{}, tc.conditions)
			assert.Equal(t, tc.expected, result)
		})
	}
}

type conditionTrue struct{}

func (c conditionTrue) Evaluate(input map[string]interface{}) (bool, error) {
	return true, nil
}

func (c conditionTrue) DebugInfo() DebugCondition {
	return DebugCondition{}
}

type conditionFalse struct{}

func (c conditionFalse) Evaluate(input map[string]interface{}) (bool, error) {
	return false, nil
}

func (c conditionFalse) DebugInfo() DebugCondition {
	return DebugCondition{}
}
