package core

import (
	"fmt"
	"github.com/JJBordy/rules/rules/functions"
)

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
	functionsDebug := make([]DebugFunction, 0)
	for _, f := range la.functions {
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
		InputPath: la.inputsPath,
		Functions: functionsDebug,
	}
}
