package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Default() map[string]Function {
	m := make(map[string]Function)

	m = defaultGeneralFunctions(m)
	m = defaultNumericFUnctions(m)
	m = defaultStringFunctions(m)

	return m
}

func defaultGeneralFunctions(m map[string]Function) map[string]Function {
	m["EMPTY"] = func(input any, args []any) (bool, error) {
		return len(fmt.Sprint(input)) == 0, nil
	}

	m["NONEMPTY"] = func(input any, args []any) (bool, error) {
		return len(fmt.Sprint(input)) > 0, nil
	}

	m["EQUAL"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("EQUAL: needs one argument")
		}
		return fmt.Sprint(input) == fmt.Sprint(args[0]), nil
	}

	return m
}

func defaultNumericFUnctions(m map[string]Function) map[string]Function {
	m["GREATER"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("GREATER", input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr > argsNr[0], nil
	}

	m["GREATER_EQ"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("GREATER_EQ", input, args, 1)
		if err != nil {
			return false, err
		}
		return inputNr >= argsNr[0], nil
	}

	m["LOWER"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("LOWER", input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr < argsNr[0], nil
	}

	m["LOWER_EQ"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("LOWER_EQ", input, args, 1)
		if err != nil {
			return false, err
		}
		return inputNr <= argsNr[0], nil
	}

	m["BETWEEN"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("BETWEEN", input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr > argsNr[0] && inputNr < argsNr[1], nil
	}

	m["BETWEEN_EQ"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("BETWEEN_EQ", input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr >= argsNr[0] && inputNr <= argsNr[1], nil
	}

	m["NOT_BETWEEN"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("NOT_BETWEEN", input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr <= argsNr[0] || inputNr >= argsNr[1], nil
	}

	m["NOT_BETWEEN_EQ"] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric("NOT_BETWEEN_EQ", input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr < argsNr[0] || inputNr > argsNr[1], nil
	}

	return m
}

func defaultStringFunctions(m map[string]Function) map[string]Function {
	m["EQUAL_IGNORE_CASE"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("EQUAL_IGNORE_CASE: needs one argument")
		}
		return strings.ToLower(fmt.Sprint(input)) == strings.ToLower(fmt.Sprint(args[0])), nil
	}

	m["EQUAL_ANY"] = func(input any, args []any) (bool, error) {
		for _, arg := range args {
			if fmt.Sprint(input) == fmt.Sprint(arg) {
				return true, nil
			}
		}
		return false, nil
	}

	m["NOT_EQUAL_ANY"] = func(input any, args []any) (bool, error) {
		for _, arg := range args {
			if fmt.Sprint(input) == fmt.Sprint(arg) {
				return false, nil
			}
		}
		return true, nil
	}

	m["STARTS_WITH"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("STARTS_WITH: needs one argument")
		}
		return strings.HasPrefix(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m["STARTS_WITH_IGNORE_CASE"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("STARTS_WITH_IGNORE_CASE: needs one argument")
		}
		return strings.HasPrefix(strings.ToLower(fmt.Sprint(input)), strings.ToLower(fmt.Sprint(args[0]))), nil
	}

	m["ENDS_WITH"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("ENDS_WITH: needs one argument")
		}
		return strings.HasSuffix(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m["ENDS_WITH_IGNORE_CASE"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("ENDS_WITH_IGNORE_CASE: needs one argument")
		}
		return strings.HasSuffix(strings.ToLower(fmt.Sprint(input)), strings.ToLower(fmt.Sprint(args[0]))), nil
	}

	m["CONTAINS"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("CONTAINS: needs one argument")
		}
		return strings.Contains(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m["CONTAINS_IGNORE_CASE"] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New("CONTAINS_IGNORE_CASE: needs one argument")
		}
		return strings.Contains(strings.ToLower(fmt.Sprint(input)), strings.ToLower(fmt.Sprint(args[0]))), nil
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
