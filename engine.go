package rules

import (
	"errors"
	"fmt"
	"strconv"
)

type Engine struct {
	conditionChains        map[string]ConditionChain
	ruleSets               map[string][]Rule
	conditionFunctions     map[string]ConditionFunc
	conditionListFunctions map[string]ConditionFuncOfList
}

func NewEngine(userFunctions map[string]ConditionFunc, userListFunctions map[string]ConditionFuncOfList) *Engine {
	e := Engine{
		ruleSets: make(map[string][]Rule),
	}

	functions := defaultFunctions()
	for k, v := range userFunctions {
		functions[k] = v
	}

	listFunctions := defaultListFunctions()
	for k, v := range userListFunctions {
		listFunctions[k] = v
	}

	e.conditionFunctions = functions
	e.conditionListFunctions = listFunctions

	return &e
}

func (e *Engine) CreateSet(setName string, ruleInputs []RuleInput) error {
	parsedRules := make([]Rule, 0)

	for _, r := range ruleInputs {
		parsedRule, err := e.parseRuleInput(r)
		if err != nil {
			return err
		}
		parsedRules = append(parsedRules, parsedRule)
	}

	e.ruleSets[setName] = parsedRules

	return nil
}

func (e *Engine) parseRuleInput(ruleInput RuleInput) (Rule, error) {
	rule := Rule{
		Name:      ruleInput.Name,
		ID:        ruleInput.ID,
		Map:       ruleInput.Map,
		Output:    ruleInput.Output,
		OutputMap: ruleInput.OutputMap,
	}

	switch ruleInput.ConditionsChain {
	case "AND":
		rule.conditionChain = ConditionChainAnd{}
	case "OR":
		rule.conditionChain = ConditionChainOr{}
	default:
		rule.conditionChain = ConditionChainAnd{}
	}

	// if OUTPUT_MAP
	// if OUTPUT_VALIDATION
	// if OUTPUT

	rule.conditions = make([]Condition, 0)
	for _, conditionInput := range ruleInput.Conditions {
		if singleInput, ok := conditionInput["input"]; ok {
			for function, args := range conditionInput {
				// pass by value for now, in the future might use a pointer to the functions, to save memory
				//e.conditionFunctions[function]
			}
		} else if listInput, ok := conditionInput["inputs"]; ok {

		} else {
			return rule, errors.New("invalid condition: no input")
		}
	}

	return rule, nil
}

func (e *Engine) EvaluateSet(setName string, input map[string]interface{}) (map[string]interface{}, error) {

	outputs := make([]map[string]interface{}, 0)

	for _, rule := range e.ruleSets[setName] {
		output, err := rule.GenerateOutput(input)
		if err != nil {
			return nil, err
		}
		if output != nil {
			outputs = append(outputs, output)
		}
	}

	return BuildOutput(outputs)
}

func defaultFunctions() map[string]ConditionFunc {
	m := make(map[string]ConditionFunc)

	m["GREATER"] = func(input any, args []any) (bool, error) {
		// will need a function that converts input and args to numbers

		inputNr, err := strconv.ParseFloat(fmt.Sprint(input), 64)
		if err != nil {
			return false, err
		}

		if len(args) == 0 {
			return false, errors.New("GREATER function requires one argument; none are provided")
		}

		argNr, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
		if err != nil {
			return false, err
		}

		return inputNr > argNr, nil
	}

	return m
}

func defaultListFunctions() map[string]ConditionFuncOfList {
	return make(map[string]ConditionFuncOfList)
}
