package functions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// All functions can have arguments, they are represented in square brackets

	// General functions

	// Empty [] - the input is empty
	Empty = "Empty"
	// NonEmpty [] - the input is not empty
	NonEmpty = "NonEmpty"
	// Equal [ any ] - the input is equal to the argument
	Equal = "Equal"

	// Numeric functions

	// Greater [nr] - the input is a number and is greater than the argument
	Greater = "Greater"
	// GreaterEq [nr] - the input is a number and is greater or equal to the argument
	GreaterEq = "GreaterEq"
	// Lower [nr] - the input is a number and is lower than the argument
	Lower = "Lower"
	// LowerEq [nr] - the input is a number and is lower or equal to the argument
	LowerEq = "LowerEq"
	// Between [nr] - the input is a number and is between the two arguments
	Between = "Between"
	// BetweenEq [nr] - the input is a number and is between or equal to the two arguments
	BetweenEq = "BetweenEq"
	// NotBetween [nr] - the input is a number and is not between the two arguments
	NotBetween = "NotBetween"
	// NotBetweenEq [nr] - the input is a number and is not between or equal to the two arguments
	NotBetweenEq = "NotBetweenEq"

	// String functions

	// EqualIgnoreCase [str] - the input is equal to the argument; case-insensitive
	EqualIgnoreCase = "EqualIgnoreCase"
	// EqualAny [str...]- the input is equal to any of the arguments
	EqualAny = "EqualAny"
	// EqualAnyIgnoreCase [str...] - the input is equal to any of the arguments; case-insensitive
	EqualAnyIgnoreCase = "EqualAnyIgnoreCase"
	// NotEqualAny [str...] - the input is not equal to any of the arguments
	NotEqualAny = "NotEqualAny"
	// StartsWith [str] - the input starts with the argument
	StartsWith = "StartsWith"
	// StartsWithIgnoreCase [str] - the input starts with the argument; case-insensitive
	StartsWithIgnoreCase = "StartsWithIgnoreCase"
	// EndsWith [str] - the input ends with the argument
	EndsWith = "EndsWith"
	// EndsWithIgnoreCase [str] - the input ends with the argument; case-insensitive
	EndsWithIgnoreCase = "EndsWithIgnoreCase"
	// Contains [str] - the input contains the argument
	Contains = "Contains"
	// ContainsIgnoreCase [str] - the input contains the argument; case-insensitive
	ContainsIgnoreCase = "ContainsIgnoreCase"
)

// Function - definition of a function, which takes one element as input
// The concrete elements of this type have the exact implementation of the function
type Function func(input any, args []any) (bool, error)

func Default() map[string]Function {
	m := make(map[string]Function)

	m = defaultGeneralFunctions(m)
	m = defaultNumericFUnctions(m)
	m = defaultStringFunctions(m)

	return m
}

func defaultGeneralFunctions(m map[string]Function) map[string]Function {
	m[Empty] = func(input any, args []any) (bool, error) {
		return len(fmt.Sprint(input)) == 0, nil
	}

	m[NonEmpty] = func(input any, args []any) (bool, error) {
		return len(fmt.Sprint(input)) > 0, nil
	}

	m[Equal] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", Equal))
		}
		return fmt.Sprint(input) == fmt.Sprint(args[0]), nil
	}

	return m
}

func defaultNumericFUnctions(m map[string]Function) map[string]Function {
	m[Greater] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(Greater, input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr > argsNr[0], nil
	}

	m[GreaterEq] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(GreaterEq, input, args, 1)
		if err != nil {
			return false, err
		}
		return inputNr >= argsNr[0], nil
	}

	m[Lower] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(Lower, input, args, 1)
		if err != nil {
			return false, err
		}

		return inputNr < argsNr[0], nil
	}

	m[LowerEq] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(LowerEq, input, args, 1)
		if err != nil {
			return false, err
		}
		return inputNr <= argsNr[0], nil
	}

	m[Between] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(Between, input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr > argsNr[0] && inputNr < argsNr[1], nil
	}

	m[BetweenEq] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(BetweenEq, input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr >= argsNr[0] && inputNr <= argsNr[1], nil
	}

	m[NotBetween] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(NotBetween, input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr <= argsNr[0] || inputNr >= argsNr[1], nil
	}

	m[NotBetweenEq] = func(input any, args []any) (bool, error) {
		inputNr, argsNr, err := ParseNumeric(NotBetweenEq, input, args, 2)
		if err != nil {
			return false, err
		}
		return inputNr < argsNr[0] || inputNr > argsNr[1], nil
	}

	return m
}

func defaultStringFunctions(m map[string]Function) map[string]Function {
	m[EqualIgnoreCase] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", EqualIgnoreCase))
		}
		return strings.EqualFold(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m[EqualAnyIgnoreCase] = func(input any, args []any) (bool, error) {
		for _, arg := range args {
			if strings.EqualFold(fmt.Sprint(input), fmt.Sprint(arg)) {
				return true, nil
			}
		}
		return false, nil
	}

	m[EqualAny] = func(input any, args []any) (bool, error) {
		for _, arg := range args {
			if fmt.Sprint(input) == fmt.Sprint(arg) {
				return true, nil
			}
		}
		return false, nil
	}

	m[NotEqualAny] = func(input any, args []any) (bool, error) {
		for _, arg := range args {
			if fmt.Sprint(input) == fmt.Sprint(arg) {
				return false, nil
			}
		}
		return true, nil
	}

	m[StartsWith] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", StartsWith))
		}
		return strings.HasPrefix(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m[StartsWithIgnoreCase] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", StartsWithIgnoreCase))
		}
		return strings.HasPrefix(strings.ToLower(fmt.Sprint(input)), strings.ToLower(fmt.Sprint(args[0]))), nil
	}

	m[EndsWith] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", EndsWith))
		}
		return strings.HasSuffix(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m[EndsWithIgnoreCase] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", EndsWithIgnoreCase))
		}
		return strings.HasSuffix(strings.ToLower(fmt.Sprint(input)), strings.ToLower(fmt.Sprint(args[0]))), nil
	}

	m[Contains] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", Contains))
		}
		return strings.Contains(fmt.Sprint(input), fmt.Sprint(args[0])), nil
	}

	m[ContainsIgnoreCase] = func(input any, args []any) (bool, error) {
		if len(args) != 1 {
			return false, errors.New(fmt.Sprintf("%s: needs one argument", ContainsIgnoreCase))
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
