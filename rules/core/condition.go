package core

import (
	"github.com/JJBordy/rules/rules/functions"
)

// Condition - a condition that can be applied to an input, it can be one of three types: single input, list input or list aggregate
type Condition interface {
	Evaluate(input map[string]any) (bool, error)
	DebugInfo() DebugCondition
}

// InputFunction - a function that can be applied to an input, contains the function itself and the arguments
// for value X, apply the function and return true if the input passes the function, false if not
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
