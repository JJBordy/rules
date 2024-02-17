package core

import (
	"errors"
	"fmt"
	"github.com/JJBordy/rules/rules/functions"
)

// Need to have condition constructor, with ConditionData inside; as a pointer, to add functions safely
// Then a Condition struct, reference by value, with the IsTrue function only

type Condition struct {
	InputPath string

	EngineFunctions           map[string]functions.Function
	EngineFunctionsOfList     map[string]functions.FunctionOfList
	ConditionFunctions        map[string][]any // function name & arguments
	ConditionsFunctionsOfList map[string][]any // function of list name & arguments
}

func NewCondition(inputPath string, engineFunctions map[string]functions.Function, engineFunctionsOfList map[string]functions.FunctionOfList) *Condition {
	return &Condition{
		InputPath:                 inputPath,
		EngineFunctions:           engineFunctions,
		EngineFunctionsOfList:     engineFunctionsOfList,
		ConditionFunctions:        make(map[string][]any),
		ConditionsFunctionsOfList: make(map[string][]any),
	}
}

// AddFunction - include function key and arguments in the condition
func (c *Condition) AddFunction(funcName string, funcArgs any) error {
	if _, isDefined := c.EngineFunctions[funcName]; isDefined {
		return c.addFunction(funcName, funcArgs)
	} else if _, isDefinedInList := c.EngineFunctionsOfList[funcName]; isDefinedInList {
		return c.addFunctionOfList(funcName, funcArgs)
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

// IsTrue - evaluate the condition for the given input
func (c *Condition) IsTrue(input map[string]interface{}) (bool, error) {
	passedFunctions := 0

	for funcKey, funcArgs := range c.ConditionFunctions {
		valid, err := c.EngineFunctions[funcKey](extractFieldVal(c.InputPath, input), funcArgs)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}

	for funcKey, funcArgs := range c.ConditionsFunctionsOfList {
		valid, err := c.EngineFunctionsOfList[funcKey](extractFromSlice(c.InputPath, input), funcArgs)
		if err != nil {
			return false, err
		}
		if valid {
			passedFunctions++
		}
	}

	return passedFunctions == (len(c.ConditionFunctions) + len(c.ConditionsFunctionsOfList)), nil
}
