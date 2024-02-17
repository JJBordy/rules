package functions

import (
	"errors"
	"fmt"
	"strconv"
)

func Default() map[string]Function {
	m := make(map[string]Function)

	m["GREATER"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("GREATER", input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr > argsNr[0], nil
	}

	m["LESS_THAN"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("LESS_THAN", input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr < argsNr[0], nil
	}

	return m
}

func ParseNumeric(functionName string, input any, args []any, requiredArgsCount int) (float64, []float64, error) {
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

	if len(argsNr) != requiredArgsCount {
		return 0, nil, errors.New(fmt.Sprintf("[%s]: not enough arguments provided", functionName))
	}

	return inputNr, argsNr, nil
}
