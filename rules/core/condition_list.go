package core

import (
	"fmt"
	"github.com/JJBordy/rules/rules/functions"
)

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
	functionsDebug := make([]DebugFunction, 0)
	for _, f := range l.functions {
		args := make([]string, 0)
		for _, arg := range f.args {
			args = append(args, fmt.Sprint(arg))
		}
		functionsDebug = append(functionsDebug, DebugFunction{
			Name: f.name,
			Args: args,
		})
	}
	return DebugCondition{
		InputPath: l.inputsPath,
		Functions: functionsDebug,
	}
}
