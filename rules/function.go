package rules

import (
	"errors"
	"fmt"
	"strconv"
)

type ConditionFuncStruct struct {
	FunctionKey string

	Func ConditionFunc
}

type ConditionFunc func(input any, args []any) (bool, error)
type ConditionFuncOfList func(inputs []any, args []any) (bool, error)

type conditionFunction struct {
	Key            string
	Args           []any
	EngineFunction func(input any, args []any) (bool, error)
}

func (c conditionFunction) Function(input any) (bool, error) {
	return c.EngineFunction(input, c.Args)
}

type conditionFunctionOfList struct {
	Key                string
	Args               []any
	EngineFunctionList func(inputs []any, args []any) (bool, error)
}

func (c conditionFunctionOfList) FunctionList(inputs []any) (bool, error) {
	return c.EngineFunctionList(inputs, c.Args)
}

func MakeAllNumeric(functionName string, input any, args []any) (float64, []float64, error) {
	inputNr, err := strconv.ParseFloat(fmt.Sprint(input), 64)
	if err != nil {
		return 0, nil,
			errors.New(fmt.Sprintf("[%s]: could not convert input [%v] to number: %s", functionName, input, err.Error()))
	}

	argsNr := make([]float64, 0)
	for i, arg := range args {
		argNr, err := strconv.ParseFloat(fmt.Sprint(arg), 64)
		if err != nil {
			return 0, nil,
				errors.New(fmt.Sprintf("[%s]: could not convert argument [%d] [%v] to number: %s", functionName, i, arg, err.Error()))
		}
		argsNr = append(argsNr, argNr)
	}

	return inputNr, argsNr, nil
}

func MakeAllListNumeric(functionName string, inputs []any, args []any) ([]float64, []float64, error) {
	inputsNr := make([]float64, 0)
	for _, input := range args {
		inputNr, err := strconv.ParseFloat(fmt.Sprint(input), 64)
		if err != nil {
			return nil, nil,
				errors.New(fmt.Sprintf("[%s]: could not convert input [%v] to number: %s", functionName, input, err.Error()))
		}
		inputsNr = append(inputsNr, inputNr)
	}

	argsNr := make([]float64, 0)
	for i, arg := range args {
		argNr, err := strconv.ParseFloat(fmt.Sprint(arg), 64)
		if err != nil {
			return nil, nil,
				errors.New(fmt.Sprintf("[%s]: could not convert argument [%d] [%v] to number: %s", functionName, i, arg, err.Error()))
		}
		argsNr = append(argsNr, argNr)
	}

	return inputsNr, argsNr, nil
}
