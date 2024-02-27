package rules

import (
	"errors"
	"fmt"
	"github.com/JJBordy/rules/rules/core"
	"github.com/JJBordy/rules/rules/functions"
	"github.com/JJBordy/rules/rules/output"
	"strconv"
)

type Engine struct {
	ruleSets map[string][]core.Rule

	conditionChains map[string]core.ConditionsChain

	singleInputFunctions map[string]functions.Function

	listAggregation functions.AggregateFunctions

	listInputConstraints map[string]functions.ListFunctionConstraint
}

type EngineConstructorData struct {
	UserFunctions map[string]functions.Function
}

func NewEngine(cd EngineConstructorData) *Engine {
	e := Engine{
		ruleSets: make(map[string][]core.Rule),
	}

	funcs := functions.Default()
	for k, v := range cd.UserFunctions {
		funcs[k] = v
	}

	e.singleInputFunctions = funcs

	e.listAggregation = functions.AllAggregateFunctions()
	e.listInputConstraints = functions.AllListFunctionConstraints()

	return &e
}

// EvaluateSet - returns the output generated by passing of the input through the set's rules
func (e *Engine) EvaluateSet(setName string, input map[string]interface{}) (map[string]interface{}, error) {

	outputs := make([]map[string]interface{}, 0)

	for _, rule := range e.ruleSets[setName] {
		out, err := rule.GenerateOutput(input)
		if err != nil {
			return nil, err
		}
		if out != nil {
			outputs = append(outputs, out)
		}
	}

	return output.BuildOutput(outputs)
}

// DebugSet - returns the output generated by passing of the input through the set's rules
// also returns a map with all the rule names and the conditions which returned false
// the key of the map map[string][]core.DebugConditions is the rule name, not the rule ID, as the ID is not guaranteed to be there
func (e *Engine) DebugSet(setName string, input map[string]interface{}) (map[string][]core.DebugConditions, map[string]interface{}, error) {
	outputs := make([]map[string]interface{}, 0)
	debugOutput := make(map[string][]core.DebugConditions)

	for _, rule := range e.ruleSets[setName] {

		debugOut, out, err := rule.DebugOutput(input)
		if err != nil {
			return nil, nil, err
		}
		if out != nil {
			outputs = append(outputs, out)
		}
		if debugOut != nil {
			debugOutput[rule.Name] = debugOut
		}
	}

	builtOutput, err := output.BuildOutput(outputs)
	if err != nil {
		return nil, nil, err
	}

	return debugOutput, builtOutput, nil
}

func (e *Engine) CreateSet(setName string, ruleInputs []RuleInput) error {
	parsedRules := make([]core.Rule, 0)

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

func (e *Engine) parseRuleInput(ruleInput RuleInput) (core.Rule, error) {
	rule := core.Rule{
		Name:       ruleInput.Name,
		ID:         ruleInput.ID,
		Map:        ruleInput.Map,
		Output:     ruleInput.Output,
		OutputMap:  ruleInput.OutputMap,
		Conditions: make([]core.Condition, 0),
	}

	conditionsChain, err := core.NewConditionChain(ruleInput.ConditionsChain)
	if err != nil {
		return rule, err
	}
	rule.ConditionChain = conditionsChain

	for _, condition := range ruleInput.Conditions {
		newCondition := core.NewCondition(fmt.Sprint(condition.Input), e.singleInputFunctions, functions.AllListFunctionConstraints())

		for function, args := range condition.Functions {
			err = newCondition.AddFunction(function, args)
			if err != nil {
				return rule, err
			}
		}

		rule.Conditions = append(rule.Conditions, *newCondition)
	}

	for _, inputListCondition := range ruleInput.ConditionsList {
		newConditionList := core.NewCondition(fmt.Sprint(inputListCondition.Inputs), e.singleInputFunctions, functions.AllListFunctionConstraints())

		// aggregation
		if inputListCondition.Aggregate.Type != "" {
			if aggregateFunction, ok := e.listAggregation.GetFunction(inputListCondition.Aggregate.Type); ok {
				newConditionList.ListAggregateType = aggregateFunction
			} else {
				return rule, errors.New(fmt.Sprintf("[RULE: %s] - unknown aggregation function: %s", ruleInput.Name, inputListCondition.Aggregate.Type))
			}

			for function, args := range inputListCondition.Aggregate.Functions {
				if _, ok := e.singleInputFunctions[function]; ok {
					newConditionList.ListAggregateFunctions[function] = args
				} else {
					return rule, errors.New(fmt.Sprintf("[RULE: %s] - unknown function: %s", ruleInput.Name, function))
				}
			}
		}

		// list functions
		if len(inputListCondition.ListFunctions.Functions) > 0 {
			newConditionList.ConditionsFunctionsOfList = make(map[string][]any)

			// list functions
			for funcName, funcArgs := range inputListCondition.ListFunctions.Functions {
				if _, ok := e.singleInputFunctions[funcName]; ok {
					newConditionList.ConditionsFunctionsOfList[funcName] = funcArgs
				} else {
					return rule, errors.New(fmt.Sprintf("[RULE: %s] - unknown function of list: %s", ruleInput.Name, funcName))
				}
			}

			// list constraints
			for listConstraint, args := range inputListCondition.ListFunctions.Constraints {
				if _, ok := e.listInputConstraints[listConstraint]; ok {
					intArgs, err := toIntArray(args)
					if err != nil {
						return rule, errors.New(fmt.Sprintf("[RULE: %s] - error parsing constraint args of %s to int: %s", ruleInput.Name, listConstraint, err))
					}
					newConditionList.ListFunctionConstraints[listConstraint] = intArgs
				} else {
					return rule, errors.New(fmt.Sprintf("[RULE: %s] - unknown constraint of list: %s", ruleInput.Name, listConstraint))
				}
			}
		}
	}

	return rule, nil
}

func toIntArray(args []any) ([]int, error) {
	intArr := make([]int, 0)
	for _, arg := range args {
		intVal, err := strconv.Atoi(fmt.Sprint(arg))
		if err != nil {
			return intArr, err
		}
		intArr = append(intArr, intVal)
	}
	return intArr, nil
}

// extractConditionInput - extracts input or inputs value, removes it from the map
func extractConditionInput(cond map[string]interface{}) (string, error) {
	if inputPath, ok := cond["input"]; ok {
		delete(cond, "input")
		return fmt.Sprint(inputPath), nil
	}
	if inputListPath, ok := cond["inputs"]; ok {
		delete(cond, "inputs")
		return fmt.Sprint(inputListPath), nil
	}
	return "", errors.New("no input specification in condition")
}
