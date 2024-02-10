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
		valid, err := conditionFunction(extractFieldVal(c.inputPath, input), c.args)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}

	for _, conditionFunctionOfList := range c.conditionFuncsList {
		valid, err := conditionFunctionOfList(extractFieldListValues(c.inputPath, input), c.args)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}

	return passedFunctions == len(c.conditionFuncs), nil
}
func extractFieldListValues(listPath string, input map[string]interface{}) []any {

	lastIsList := false
	if strings.HasSuffix(listPath, "[*]") {
		lastIsList = true
	}
	lastI := len(strings.Split(listPath, "[*]")) - 1

	//resultsArray := make([]any, 0)
	pathsToLists := strings.Split(listPath, "[*]")
	for i, pathToList := range pathsToLists {
		workMap := input
		for _, fieldName := range strings.Split(pathToList, ".") {
			if val, ok := workMap[fieldName].(map[string]interface{}); ok {
				workMap = val
			} else if arr, ok := workMap[fieldName].([]any); ok {
				//resultsArray = append(resultsArray, extractAsList(fieldName, arr))
				return arr
			} else {
				return []any{workMap[fieldName]}
			}
		}
		if i == lastI && lastIsList {
			// process the last one as a list
		}
	}

	return nil
}

type ExtractResult struct {
	List  []any
	Value string
}

func extractAsList(path string, input map[string]interface{}) []any {

	fullPath := strings.Split(path, ".")
	resultingArray := make([]any, 0)
	workMap := input

	for pi, pathElem := range fullPath {
		if mp, ok := workMap[pathElem].(map[string]interface{}); ok {
			workMap = mp
		} else if arr, ok := workMap[pathElem].([]map[string]interface{}); ok {
			for _, arrElem := range arr {
				resultingArray = append(resultingArray, extractAsList(strings.Join(fullPath[pi+1:], "."), arrElem)...)
			}
		} else {
			resultingArray = append(resultingArray, workMap[pathElem])
		}
	}

	cleanedArray := make([]any, 0)
	for _, elem := range resultingArray {
		if elem != nil {
			cleanedArray = append(cleanedArray, elem)
		}
	}

	return cleanedArray
}

func extractFieldVal(path string, input map[string]interface{}) string {
	workMap := input
	for _, fieldName := range strings.Split(path, ".") {
		if innerMap, ok := workMap[fieldName].(map[string]interface{}); ok {
			workMap = innerMap
		} else {
			return fmt.Sprint(workMap[fieldName])
		}
	}

	return ""
}

func (c Condition) Validate(input map[string]interface{}) error {
	return nil
}
