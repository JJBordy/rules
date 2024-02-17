package core

import "github.com/JJBordy/rules/rules/functions"

/* REDUNTANT, can call directly the engine's funtions; Condition can keep function key and args */

// RuleFunction - a function contained in a condition of rule
// functions.Function implementations are defined in (go) code, while RuleFunction is read from the rule (yaml)
type RuleFunction struct {
	Key      string
	Args     []any
	Function functions.Function
}

// Execute - executes the contained functions.Function with the arguments specified in the rule
func (c RuleFunction) Execute(input any) (bool, error) {
	return c.Function(input, c.Args)
}

// RuleFunctionOfList - a function contained in a condition of rule
// functions.FunctionOfList implementations are defined in (go) code, while RuleFunction is read from the rule (yaml)
type RuleFunctionOfList struct {
	Key            string
	Args           []any
	FunctionOfList functions.FunctionOfList
}

// ExecuteOfList - executes the contained functions.ExecuteOfList with the arguments specified in the rule
func (c RuleFunctionOfList) ExecuteOfList(inputs []any) (bool, error) {
	return c.FunctionOfList(inputs, c.Args)
}
