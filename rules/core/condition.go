package core

import (
	"errors"
	"fmt"
	"github.com/JJBordy/rules/rules/functions"
)

type ConditionI interface {
	Evaluate(input map[string]any) (bool, error)
}

// InputFunction - a function that can be applied to an input, contains the function itself and the arguments
type InputFunction struct {
	name      string
	args      []any
	execution functions.Function
}

func (f InputFunction) ExecuteFunction(input any) (bool, error) {
	return f.execution(input, f.args)
}

type SingleInputCondition struct {
	inputPath string
	functions []InputFunction
}

func (s SingleInputCondition) Evaluate(input map[string]any) (bool, error) {
	for _, f := range s.functions {
		valid, err := f.ExecuteFunction(extractFieldVal(s.inputPath, input))
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}

type ListInputConstraint struct {
	args       []int
	constraint functions.ListFunctionConstraint
}

func (lc ListInputConstraint) ExecuteConstraint(listLength int, validResults int) bool {
	return lc.constraint(listLength, validResults, lc.args)
}

type ListInputCondition struct {
	inputsPath  string
	functions   []InputFunction
	constraints []ListInputConstraint
}

func (l ListInputCondition) Evaluate(input map[string]any) (bool, error) {
	listElements := extractFromSlice(l.inputsPath, input)
	listValidResults := 0

	for _, listElem := range listElements { // evaluating if each element of the list passes all the functions
		elemPassedFunctions := 0

		for _, f := range l.functions {
			passed, err := f.ExecuteFunction(listElem)
			if err != nil {
				return false, err
			}
			if passed {
				elemPassedFunctions++
			}
		}

		if elemPassedFunctions == len(l.functions) {
			listValidResults++
		}
	}

	for _, c := range l.constraints {
		constraintResult := c.ExecuteConstraint(len(listElements), listValidResults)
		if !constraintResult {
			return false, nil
		}
	}

	return true, nil
}

type ListAggregateCondition struct {
	inputsPath string
	functions  []InputFunction

	aggregate functions.AggregateFunction
}

func (la ListAggregateCondition) Evaluate(input map[string]any) (bool, error) {

	listElements := extractFromSlice(la.inputsPath, input)

	aggregationResult, err := la.aggregate(listElements)
	if err != nil {
		return false, err
	}

	for _, f := range la.functions {
		valid, err := f.ExecuteFunction(aggregationResult)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}

	return true, nil
}

// Need to have condition constructor, with ConditionData inside; as a pointer, to add functions safely
// Then a Condition struct, reference by value, with the Evaluate function only

type Condition struct {
	InputPath string

	// functions defined in the engine
	EngineFunctions map[string]functions.Function

	// definitions of functions to apply for the condition to be valid
	// single input
	ConditionFunctions map[string][]any // function name & arguments
	// list input
	ConditionsFunctionsOfList     map[string][]any
	EngineListFunctionConstraints map[string]functions.ListFunctionConstraint

	ListFunctionConstraints map[string][]int // constraint name & arguments

	ListAggregateType      functions.AggregateFunction
	ListAggregateFunctions map[string][]any
}

func NewCondition(inputPath string, engineFunctions map[string]functions.Function, engineListFunctionConstraints map[string]functions.ListFunctionConstraint) *Condition {
	return &Condition{
		InputPath:                     inputPath,
		EngineFunctions:               engineFunctions,
		ConditionFunctions:            make(map[string][]any),
		ConditionsFunctionsOfList:     make(map[string][]any),
		EngineListFunctionConstraints: engineListFunctionConstraints,
	}
}

// AddFunction - include function key and arguments in the condition
func (c *Condition) AddFunction(funcName string, funcArgs any) error {
	if _, isDefined := c.EngineFunctions[funcName]; isDefined {
		return c.addFunction(funcName, funcArgs)
	} else {
		return errors.New(fmt.Sprintf("function: [%s] is not defined in the engine", funcName))
	}
}

func (c *Condition) addFunction(funcName string, funcArgs any) error {
	if args, isSlice := funcArgs.([]any); isSlice {
		c.ConditionFunctions[funcName] = args
		return nil
	}
	return errors.New(fmt.Sprintf("function: [%s] arguments must be of type []any", funcName))
}

func (c *Condition) addFunctionOfList(funcName string, funcArgs any) error {
	if args, ok := funcArgs.([]any); ok {
		c.ConditionsFunctionsOfList[funcName] = args
	}
	return errors.New(fmt.Sprintf("function: [%s] arguments must be of type []any", funcName))
}

// Evaluate - evaluate the condition for the given input
func (c *Condition) Evaluate(input map[string]interface{}) (bool, error) {

	// single input evaluation
	for funcName, funcArgs := range c.ConditionFunctions {
		valid, err := c.EngineFunctions[funcName](extractFieldVal(c.InputPath, input), funcArgs)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}

	// list aggregate evaluation
	if c.ListAggregateType != nil {
		aggregationResult, err := c.ListAggregateType(extractFromSlice(c.InputPath, input))
		if err != nil {
			return false, err
		}
		for funcName, funcArgs := range c.ListAggregateFunctions {
			resultTrue, err := c.EngineFunctions[funcName](aggregationResult, funcArgs)
			if err != nil {
				return false, err
			}
			if !resultTrue {
				return false, nil
			}
		}
	}

	// functions for lists

	listElements := extractFromSlice(c.InputPath, input)
	listValidResults := 0

	for _, listElem := range listElements {

		elemPassedFunctions := 0

		for funcName, funcArgs := range c.ConditionsFunctionsOfList {

			passed, err := c.EngineFunctions[funcName](listElem, funcArgs)
			if err != nil {
				return false, err
			}
			if passed {
				elemPassedFunctions++
			}
		}

		if elemPassedFunctions == len(c.ConditionsFunctionsOfList) {
			listValidResults++
		}

	}

	for listFunctionConstraint, constraintArgs := range c.ListFunctionConstraints {
		constraintResult := c.EngineListFunctionConstraints[listFunctionConstraint](len(listElements), listValidResults, constraintArgs)
		if !constraintResult {
			return false, nil
		}
	}

	return true, nil
}
