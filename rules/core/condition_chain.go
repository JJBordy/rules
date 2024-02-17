package core

import (
	"errors"
	"fmt"
)

const (
	ChainTypeAND = "AND"
	ChainTypeOR  = "OR"
)

var validChainTypes = map[string]bool{
	ChainTypeAND: true,
	ChainTypeOR:  true,
}

type ConditionsChain struct {
	doDebug             bool
	debugOutput         []DebugConditions
	conditionsChainType string
}

// should have 'set debug'

func NewConditionChain(conditionsChainType string) (*ConditionsChain, error) {

	if conditionsChainType != "" && validChainTypes[conditionsChainType] == false {
		return nil, errors.New(fmt.Sprintf("invalid chain type: [%s]; valid types: %v", conditionsChainType, validChainTypes))
	}

	return &ConditionsChain{
		debugOutput:         make([]DebugConditions, 0),
		conditionsChainType: conditionsChainType,
	}, nil
}

func (cc *ConditionsChain) TurnDebugON() {
	cc.doDebug = true
}

func (cc *ConditionsChain) TurnDebugOFF() {
	cc.doDebug = false
}

func (cc *ConditionsChain) EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error) {
	switch cc.conditionsChainType {
	case ChainTypeAND:
		return cc.evaluateAND(input, conditions)
	case ChainTypeOR:
		return cc.evaluateOR(input, conditions)
	default:
		return cc.evaluateAND(input, conditions)
	}
}

func (cc *ConditionsChain) evaluateAND(input map[string]interface{}, conditions []Condition) (bool, error) {
	passedConditions := 0

	for _, condition := range conditions {
		passed, err := cc.evaluateCondition(input, condition)
		if err != nil {
			return false, err
		}
		if passed {
			passedConditions++
		}
	}

	return passedConditions == len(conditions), nil
}

func (cc *ConditionsChain) evaluateOR(input map[string]interface{}, conditions []Condition) (bool, error) {
	for _, condition := range conditions {
		passed, err := cc.evaluateCondition(input, condition)
		if err != nil {
			return false, err
		}
		if passed {
			return true, nil
		}
	}

	return false, nil
}

func (cc *ConditionsChain) DebugOutput() []DebugConditions {
	return cc.debugOutput
}

func (cc *ConditionsChain) evaluateCondition(input map[string]interface{}, c Condition) (bool, error) {

	defer func() {
		if cc.doDebug {
			cc.debugCondition(c)
		}
	}()

	return c.IsTrue(input)
}

type DebugConditions struct {
	Input            string
	FailedConditions []DebugCondition
}

type DebugCondition struct {
	FunctionName string
	FunctionArgs []any
}

func (cc *ConditionsChain) debugCondition(c Condition) {
	debugConditions := make([]DebugCondition, 0)

	for funcName, funcArgs := range c.ConditionFunctions {
		debugConditions = append(debugConditions, DebugCondition{
			FunctionName: funcName,
			FunctionArgs: funcArgs,
		})
	}

	for funcName, funcArgs := range c.ConditionsFunctionsOfList {
		debugConditions = append(debugConditions, DebugCondition{
			FunctionName: funcName,
			FunctionArgs: funcArgs,
		})
	}

	cc.debugOutput = append(cc.debugOutput, DebugConditions{
		Input:            c.InputPath,
		FailedConditions: debugConditions,
	})
}
