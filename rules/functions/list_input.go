package functions

import (
	"errors"
	"fmt"
	"strconv"
)

func DefaultList() map[string]FunctionOfList {
	return make(map[string]FunctionOfList)
}

func ParseListNumeric(functionName string, inputs []any, args []any, requiredArgsCount int) ([]float64, []float64, error) {
	inputsNr := make([]float64, 0)
	for _, input := range inputs {
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

	if len(argsNr) != requiredArgsCount {
		return nil, nil, errors.New(fmt.Sprintf("[%s]: not enough arguments provided", functionName))
	}

	return inputsNr, argsNr, nil
}
