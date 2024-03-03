package core

import (
	"errors"
	"fmt"
)

const (
	// AND - all conditions should be true for the rule to produce the output (default if none specified)
	AND = "AND"
	// OR - at least one condition should be true for the rule to pass
	OR = "OR"
	// NAND - if all conditions are true, the rule will not produce the output
	NAND = "NAND"
	// NOR - if at least one condition is true, the rule will not pass
	NOR = "NOR"
	// XOR - if exactly one condition is true, the rule will pass
	XOR = "XOR"
	// XNOR - if exactly one condition is true, the rule will not pass
	XNOR = "XNOR"
)

var validChainTypes = map[string]bool{
	AND:  true,
	OR:   true,
	NAND: true,
	NOR:  true,
	XOR:  true,
	XNOR: true,
}

// ConditionsChain - all rules are a set of conditions
// those conditions are tied by a chain; the default one is 'AND' - all conditions must pass for the rule to pass and give its output
// but the user can select other types of chains, in the form of the other logical operators
// the condition chain, when evaluating the conditions, can also return a debug output, indicating which conditions passed or failed
type ConditionsChain struct {
	doDebug             bool
	debugOutput         []DebugCondition
	conditionsChainType string
}

func NewConditionChain(conditionsChainType string) (*ConditionsChain, error) {

	if conditionsChainType != "" && validChainTypes[conditionsChainType] == false {
		return nil, errors.New(fmt.Sprintf("invalid chain type: [%s]; valid types: %v", conditionsChainType, validChainTypes))
	}

	return &ConditionsChain{
		debugOutput:         make([]DebugCondition, 0),
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
	case AND:
		return cc.evaluateAND(input, conditions)
	case OR:
		return cc.evaluateOR(input, conditions)
	case NAND:
		return cc.evaluateNAND(input, conditions)
	case NOR:
		return cc.evaluateNOR(input, conditions)
	case XOR:
		return cc.evaluateXOR(input, conditions)
	case XNOR:
		return cc.evaluateXNOR(input, conditions)
	default:
		return cc.evaluateAND(input, conditions)
	}
}

func (cc *ConditionsChain) evaluateXNOR(input map[string]interface{}, conditions []Condition) (bool, error) {
	passed := 0
	for _, c := range conditions {
		passedCondition, err := cc.evaluateCondition(input, c)
		if err != nil {
			return false, err
		}
		if passedCondition {
			passed++
		}
	}
	if passed == 1 {
		return false, nil
	}

	return true, nil
}

func (cc *ConditionsChain) evaluateXOR(input map[string]interface{}, conditions []Condition) (bool, error) {
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

	return passedConditions == 1, nil
}

func (cc *ConditionsChain) evaluateNOR(input map[string]interface{}, conditions []Condition) (bool, error) {
	passed, err := cc.evaluateOR(input, conditions)
	return !passed, err
}

func (cc *ConditionsChain) evaluateNAND(input map[string]interface{}, conditions []Condition) (bool, error) {
	passed, err := cc.evaluateAND(input, conditions)
	return !passed, err
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

func (cc *ConditionsChain) DebugOutput() []DebugCondition {
	return cc.debugOutput
}

func (cc *ConditionsChain) evaluateCondition(input map[string]interface{}, c Condition) (bool, error) {

	defer func() {
		if cc.doDebug {
			cc.debugCondition(c)
		}
	}()

	return c.Evaluate(input)
}

type DebugConditions struct {
	Input            string
	FailedConditions []DebugCondition
}

type DebugCondition struct {
	InputPath string
	Functions []DebugFunction
}

type DebugFunction struct {
	Name string
	Args []string
}

func (cc *ConditionsChain) debugCondition(c Condition) {
	cc.debugOutput = append(cc.debugOutput, c.DebugInfo())
}
