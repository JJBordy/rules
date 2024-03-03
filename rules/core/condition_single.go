package core

import "fmt"

var _ Condition = SingleInputCondition{}

// SingleInputCondition - a condition that can be applied to a single input
type SingleInputCondition struct {
	inputPath string
	functions []InputFunction
}

func NewSingleInputCondition(inputPath string, functions []InputFunction) SingleInputCondition {
	return SingleInputCondition{
		inputPath: inputPath,
		functions: functions,
	}
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
	functionsDebug := make([]DebugFunction, 0)
	for _, f := range s.functions {
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
		InputPath: s.inputPath,
		Functions: functionsDebug,
	}
}
