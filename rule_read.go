package rules

import (
	"fmt"
	"strings"
)

type RuleInput struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id"`
	Priority int    `yaml:"priority"`

	ConditionsChain   string                   `yaml:"COND_CHAIN"`
	ConditionsMinimum int                      `yaml:"COND_MIN"`
	Conditions        []map[string]interface{} `yaml:"COND"`

	Output           map[string]interface{} `yaml:"OUTPUT"`
	OutputValidation string                 `yaml:"OUTPUT_VALIDATION"`

	Map       map[string]interface{} `yaml:"MAP"`
	OutputMap map[string]string      `yaml:"OUTPUT_MAP"`
}

type ConditionFunc func(input any, args []any) (bool, error)
type ConditionFuncOfList func(inputs []any, args []any) (bool, error)

type Rule struct {
	Name     string
	ID       string
	Priority int

	conditions []Condition

	conditionChain ConditionChain

	Map map[string]interface{}

	Output    map[string]interface{}
	OutputMap map[string]interface{}
}

func ParseRuleInput(ruleInput RuleInput) (Rule, error) {
	return Rule{}, nil
}

func (r Rule) GenerateOutput(input map[string]interface{}) (map[string]interface{}, error) {
	rulePasses, err := r.conditionChain.EvaluateConditions(input, r.conditions)
	if err != nil {
		return nil, err
	}
	if rulePasses {
		// if outputMap present, merge it and return it too
		// also, output from validation
		return r.Output, nil
	}
	return nil, nil
}

// ConditionChain - AND, OR, XAND, NOR, etc
type ConditionChain interface { // may have condition chain type - as interface
	EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error)
}

func ConditionChainAnd(input map[string]interface{}, conditions []Condition) (bool, error) {
	passedConditions := 0

	for _, condition := range conditions {
		passed, err := condition.Evaluate(input)
		if err != nil {
			return false, err
		}
		if passed {
			passedConditions++
		}
	}

	return passedConditions == len(conditions), nil
}

type Condition struct {
	inputPath          string
	args               []any
	conditionFuncs     []ConditionFunc
	conditionFuncsList []ConditionFuncOfList
}

func (c Condition) Evaluate(input map[string]interface{}) (bool, error) {
	passedFunctions := 0

	for _, conditionFunction := range c.conditionFuncs {
		// condition func for list vs condition func for single item
		valid, err := conditionFunction(extractFieldVal(c.inputPath, input), c.args)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}
	return passedFunctions == len(c.conditionFuncs), nil
}

func extractFieldVal(path string, input map[string]interface{}) string {
	// TODO: will need another one to extract from arrays
	workMap := input
	for _, fieldName := range strings.Split(path, ".") {
		if val, ok := workMap[fieldName].(map[string]interface{}); ok {
			workMap = val
		} else {
			return fmt.Sprint(workMap[fieldName])
		}
	}

	return ""
}
