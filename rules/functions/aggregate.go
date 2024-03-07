package functions

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	// MIN - minimum value in list
	MIN = "MIN"
	// MAX - maximum value in list
	MAX = "MAX"
	// SUM - sum of values in list
	SUM = "SUM"
	// AVG - average of values in list
	AVG = "AVG"
	// COUNT - count of values in list
	COUNT = "COUNT"
)

type AggregateFunction func(inputs []any) (float64, error)

type AggregateFunctions struct {
	aggregateFunctions map[string]AggregateFunction
}

func (af AggregateFunctions) GetFunction(functionName string) (AggregateFunction, bool) {
	f, ok := af.aggregateFunctions[functionName]
	return f, ok
}

func AllAggregateFunctions() AggregateFunctions {
	return AggregateFunctions{
		aggregateFunctions: map[string]AggregateFunction{
			MIN:   aggregateMIN(),
			MAX:   aggregateMAX(),
			SUM:   aggregateSUM(),
			AVG:   aggregateAVG(),
			COUNT: aggregateCOUNT(),
		},
	}
}

func aggregateCOUNT() AggregateFunction {
	return func(inputs []any) (float64, error) {
		return float64(len(inputs)), nil
	}
}

func aggregateSUM() AggregateFunction {
	return func(inputs []any) (float64, error) {
		inputsAsNumbers, err := parseInputsToNumbers(inputs)
		if err != nil {
			return 0, errors.New(fmt.Sprint("aggregate SUM: ", err.Error()))
		}

		sum := 0.0
		for _, nr := range inputsAsNumbers {
			sum += nr
		}

		return sum, nil
	}
}

func aggregateAVG() AggregateFunction {
	return func(inputs []any) (float64, error) {
		sum, err := aggregateSUM()(inputs)
		if err != nil {
			return 0, err
		}

		return sum / float64(len(inputs)), nil
	}
}

func aggregateMIN() AggregateFunction {
	return func(inputs []any) (float64, error) {
		inputsAsNumbers, err := parseInputsToNumbers(inputs)
		if err != nil {
			return 0, errors.New(fmt.Sprint("aggregate MIN: ", err.Error()))
		}

		minNr := math.MaxFloat64
		for _, nr := range inputsAsNumbers {
			if nr < minNr {
				minNr = nr
			}
		}

		return minNr, nil
	}
}

func aggregateMAX() AggregateFunction {
	return func(inputs []any) (float64, error) {
		inputsAsNumbers, err := parseInputsToNumbers(inputs)
		if err != nil {
			return 0, errors.New(fmt.Sprint("aggregate MAX: ", err.Error()))
		}

		maxNr := -math.MaxFloat64
		for _, nr := range inputsAsNumbers {
			if nr > maxNr {
				maxNr = nr
			}
		}

		return maxNr, nil
	}
}

func parseInputsToNumbers(inputs []any) ([]float64, error) {
	if len(inputs) == 0 {
		return []float64{}, errors.New("no inputs provided")
	}

	inputsNr := make([]float64, 0)
	for _, input := range inputs {
		inputNr, err := strconv.ParseFloat(fmt.Sprint(input), 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("could not convert input [%v] to number: %s", input, err.Error()))
		}
		inputsNr = append(inputsNr, inputNr)
	}

	return inputsNr, nil
}
