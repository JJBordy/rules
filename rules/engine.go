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
	ruleSets map[string][]core.RuleNew

	conditionChains map[string]core.ConditionsChain

	inputFunctions       map[string]functions.Function
	listAggregation      functions.AggregateFunctions
	listInputConstraints map[string]functions.ListFunctionConstraint
}

type EngineConstructorData struct {
	UserFunctions map[string]functions.Function
}

func NewEngine(cd EngineConstructorData) *Engine {
	e := Engine{
		ruleSets: make(map[string][]core.RuleNew),
	}

	funcs := functions.Default()
	for k, v := range cd.UserFunctions {
		funcs[k] = v
	}

	e.inputFunctions = funcs

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

func (e *Engine) CreateSet(setName string, ruleInputs []RuleInputNew) error {
	parsedRules := make([]core.RuleNew, 0)

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

func (e *Engine) parseRuleInput(ruleInput RuleInputNew) (core.RuleNew, error) {
	rule := core.RuleNew{
		Name:       ruleInput.Name,
		Map:        ruleInput.Map,
		Output:     ruleInput.Output,
		OutputMap:  ruleInput.OutputMap,
		Conditions: make([]core.ConditionI, 0),
	}

	conditionsChain, err := core.NewConditionChain(ruleInput.ConditionsChainType)
	if err != nil {
		return rule, err
	}
	rule.ConditionChain = conditionsChain

	// single input conditions
	for _, sic := range ruleInput.Conditions.SingleInputConditions {
		inputFunctions := make([]core.InputFunction, 0)

		for funcName, funcArgs := range sic.Functions {
			if function, ok := e.inputFunctions[funcName]; ok {
				inputFunctions = append(inputFunctions, core.NewInputFunction(funcName, funcArgs, function))
			} else {
				return rule, errors.New(fmt.Sprintf("[RULE: %s] function: [%s: [%+v]] is not defined in the engine", ruleInput.Name, funcName, funcArgs))
			}
		}

		singleInputCondition := core.NewSingleInputCondition(sic.InputPath, inputFunctions)
		rule.Conditions = append(rule.Conditions, singleInputCondition)
	}

	// list aggregate functions
	for _, lac := range ruleInput.Conditions.ListAggregateConditions {

		inputFunctions := make([]core.InputFunction, 0)
		for funcName, funcArgs := range lac.Functions {
			if function, ok := e.inputFunctions[funcName]; ok {
				inputFunctions = append(inputFunctions, core.NewInputFunction(funcName, funcArgs, function))
			} else {
				return rule, errors.New(fmt.Sprintf("[RULE: %s] function: [%s: [%+v]] is not defined in the engine", ruleInput.Name, funcName, funcArgs))
			}
		}

		if aggregateFunction, ok := e.listAggregation.GetFunction(lac.AggregateType); ok {
			aggregateCondition := core.NewAggregateCondition(lac.InputsPath, inputFunctions, aggregateFunction)
			rule.Conditions = append(rule.Conditions, aggregateCondition)
		} else {
			return rule, errors.New(fmt.Sprintf("[RULE: %s] - unknown aggregation function: %s", ruleInput.Name, lac.AggregateType))
		}
	}

	// list input conditions
	for _, lic := range ruleInput.Conditions.ListInputConditions {

		inputFunctions := make([]core.InputFunction, 0)
		for funcName, funcArgs := range lic.Functions {
			if function, ok := e.inputFunctions[funcName]; ok {
				inputFunctions = append(inputFunctions, core.NewInputFunction(funcName, funcArgs, function))
			} else {
				return rule, errors.New(fmt.Sprintf("[RULE: %s] function: [%s: [%+v]] is not defined in the engine", ruleInput.Name, funcName, funcArgs))
			}
		}

		listInputConstraints := make([]core.ListInputConstraint, 0)
		for constraintName, constraintArgs := range lic.Constraints {
			if listConstraintFunction, ok := e.listInputConstraints[constraintName]; ok {
				argsAsInt := make([]int, 0)
				for _, arg := range constraintArgs {
					if i, err := strconv.Atoi(fmt.Sprint(arg)); err == nil {
						argsAsInt = append(argsAsInt, i)
					} else {
						return rule, errors.New(fmt.Sprintf("[RULE: %s] constraint: [%s: [%+v]] args are not integers", ruleInput.Name, constraintName, constraintArgs))
					}
				}

				listInputConstraints = append(listInputConstraints, core.NewListInputConstraint(argsAsInt, listConstraintFunction))
			} else {
				return rule, errors.New(fmt.Sprintf("[RULE: %s] constraint: [%s: [%+v]] is not defined in the engine", ruleInput.Name, constraintName, constraintArgs))
			}
		}

		listInputCondition := core.NewListInputCondition(lic.InputsPath, inputFunctions, listInputConstraints)
		rule.Conditions = append(rule.Conditions, listInputCondition)
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
