package core

import (
	"fmt"
	"github.com/JJBordy/rules/rules/functions"
)

type ConditionI interface {
	Evaluate(input map[string]any) (bool, error)
	DebugInfo() DebugCondition
}

// InputFunction - a function that can be applied to an input, contains the function itself and the arguments
type InputFunction struct {
	name      string
	args      []any
	execution functions.Function
}

func NewInputFunction(name string, args []any, execution functions.Function) InputFunction {
	return InputFunction{
		name:      name,
		args:      args,
		execution: execution,
	}
}

func (f InputFunction) ExecuteFunction(input any) (bool, error) {
	return f.execution(input, f.args)
}

// SingleInputCondition - a condition that can be applied to a single input
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

func (s SingleInputCondition) DebugInfo() DebugCondition {
	functionsPrint := make([]string, 0)
	for _, f := range s.functions {
		functionsPrint = append(functionsPrint, fmt.Sprintf("%#v", f))
	}
	return DebugCondition{
		InputPath: s.inputPath,
		Functions: functionsPrint,
	}
}

func NewSingleInputCondition(inputPath string, functions []InputFunction) SingleInputCondition {
	return SingleInputCondition{
		inputPath: inputPath,
		functions: functions,
	}
}

type ListInputConstraint struct {
	args       []int
	constraint functions.ListFunctionConstraint
}

func NewListInputConstraint(args []int, constraint functions.ListFunctionConstraint) ListInputConstraint {
	// TODO: validate number of arguments passed to the function (1 or 2, when creating a new one)
	return ListInputConstraint{
		args:       args,
		constraint: constraint,
	}
}

func (lc ListInputConstraint) ExecuteConstraint(listLength int, validResults int) bool {
	return lc.constraint(listLength, validResults, lc.args)
}

type ListInputCondition struct {
	inputsPath  string
	functions   []InputFunction
	constraints []ListInputConstraint
}

func NewListInputCondition(inputsPath string, functions []InputFunction, constraints []ListInputConstraint) ListInputCondition {
	return ListInputCondition{
		inputsPath:  inputsPath,
		functions:   functions,
		constraints: constraints,
	}
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

func (l ListInputCondition) DebugInfo() DebugCondition {
	functionsPrint := make([]string, 0)
	for _, c := range l.constraints {
		functionsPrint = append(functionsPrint, fmt.Sprintf("%#v", c))
	}
	for _, f := range l.functions {
		functionsPrint = append(functionsPrint, fmt.Sprintf("%#v", f))
	}
	return DebugCondition{
		InputPath: l.inputsPath,
		Functions: functionsPrint,
	}
}

type ListAggregateCondition struct {
	inputsPath string
	functions  []InputFunction

	aggregate functions.AggregateFunction
}

func NewAggregateCondition(inputsPath string, functions []InputFunction, aggregate functions.AggregateFunction) ListAggregateCondition {
	return ListAggregateCondition{
		inputsPath: inputsPath,
		functions:  functions,
		aggregate:  aggregate,
	}
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

func (la ListAggregateCondition) DebugInfo() DebugCondition {
	functionsPrint := make([]string, 0)

	for _, f := range la.functions {
		functionsPrint = append(functionsPrint, fmt.Sprintf("%#v", f))
	}
	return DebugCondition{
		InputPath: la.inputsPath,
		Functions: functionsPrint,
	}
}
