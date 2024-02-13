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

type EngineConstructorData struct {
	UserFunctions     map[string]ConditionFunc
	UserListFunctions map[string]ConditionFuncOfList
}

func NewEngine(cd EngineConstructorData) *Engine {
	e := Engine{
		ruleSets: make(map[string][]Rule),
	}

	functions := defaultFunctions()
	for k, v := range cd.UserFunctions {
		functions[k] = v
	}

	listFunctions := defaultListFunctions()
	for k, v := range cd.UserListFunctions {
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

	if rule.Output != nil {

		rule.conditions = make([]Condition, 0)

		for _, conditionInput := range ruleInput.Conditions {

			if singleInputPath, ok := conditionInput["input"]; ok {

				newCondition := Condition{
					inputPath:      fmt.Sprint(singleInputPath),
					conditionFuncs: make([]conditionFunction, 0),
				}
				for function, args := range conditionInput {
					if function == "input" {
						continue
					}
					if engineFunction, ok := e.conditionFunctions[function]; ok {
						argsVal := make([]any, 0)
						for _, arg := range args.([]any) {
							argsVal = append(argsVal, arg)
						}
						newCondition.conditionFuncs = append(newCondition.conditionFuncs, conditionFunction{
							Key:            function,
							Args:           argsVal,
							EngineFunction: engineFunction,
						})
					} else {
						return rule, errors.New(fmt.Sprintf("Rule: [%s]; No such function: [%s]", rule.Name, function))
					}
				}

				rule.conditions = append(rule.conditions, newCondition)

			} else if listInputsPath, ok := conditionInput["inputs"]; ok {
				newCondition := Condition{
					inputPath:      fmt.Sprint(listInputsPath),
					conditionFuncs: make([]conditionFunction, 0),
				}
				for function, args := range conditionInput {
					if function == "inputs" {
						continue
					}
					if engineFunction, ok := e.conditionListFunctions[function]; ok {
						argsVal := make([]any, 0)
						for _, arg := range args.([]any) {
							argsVal = append(argsVal, arg)
						}
						newCondition.conditionFuncsList = append(newCondition.conditionFuncsList, conditionFunctionOfList{
							Key:                function,
							Args:               argsVal,
							EngineFunctionList: engineFunction,
						})
					}
				}
				rule.conditions = append(rule.conditions, newCondition)
			} else {
				return rule, errors.New("invalid condition: no input/s")
			}
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

	m["LESS_THAN"] = func(input any, args []any) (bool, error) {
		// will need a function that converts input and args to numbers

		inputNr, err := strconv.ParseFloat(fmt.Sprint(input), 64)
		if err != nil {
			return false, err
		}

		if len(args) == 0 {
			return false, errors.New("LESS_THAN function requires one argument; none are provided")
		}

		argNr, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
		if err != nil {
			return false, err
		}

		return inputNr < argNr, nil
	}

	return m
}

func defaultListFunctions() map[string]ConditionFuncOfList {
	return make(map[string]ConditionFuncOfList)
}
