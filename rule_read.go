package rules

import (
	"strings"
)

type RuleInput struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id"`
	Priority int    `yaml:"priority"`

	ConditionsChain string                   `yaml:"COND_CHAIN"`
	Conditions      []map[string]interface{} `yaml:"COND"`

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
	OutputMap map[string]string
}

func (r Rule) GenerateOutput(input map[string]interface{}) (map[string]interface{}, error) {
	// if OUTPUT_MAP
	// if OUTPUT_VALIDATION
	// if OUTPUT

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
		valid, err := conditionFunctionOfList(extractFromSlice(c.inputPath, input), c.args)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}

	return passedFunctions == len(c.conditionFuncs), nil
}

// extracts the values from input which contains lists
// car.windows[*].safety.ratings[*].certification will return certifications of all ratings of all windows of the car
func extractFromSlice(path string, input map[string]interface{}) []any {

	pathElems := strings.Split(path, ".")
	resultSlice := make([]any, 0)
	workMap := input

	slicePath := ""

	for pi, pathElem := range pathElems {
		if strings.HasSuffix(pathElem, "[*]") {

			slicePath = strings.TrimSuffix(pathElem, "[*]")

			if arr, ok := workMap[slicePath].([]map[string]interface{}); ok {
				for _, arrElem := range arr {
					resultSlice = append(resultSlice, extractFromSlice(strings.Join(pathElems[pi+1:], "."), arrElem)...)
				}
			}

		} else if mp, ok := workMap[pathElem].(map[string]interface{}); ok {
			workMap = mp
		} else {
			if singularVal, ok := workMap[pathElem]; ok {
				resultSlice = append(resultSlice, singularVal)
			}
		}
	}

	return resultSlice
}

// extracts the value from input, specified by the nested path, separated by "."
// for example: car.trunk.color
func extractFieldVal(path string, input map[string]interface{}) any {
	workMap := input
	for _, fieldName := range strings.Split(path, ".") {
		if innerMap, ok := workMap[fieldName].(map[string]interface{}); ok {
			workMap = innerMap
		} else {
			return workMap[fieldName]
		}
	}

	return nil
}
