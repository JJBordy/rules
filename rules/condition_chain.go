package rules

// ConditionChain - AND, OR, XAND, NOR, etc
type ConditionChain interface { // may have condition chain type - as interface
	EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error)
}

type ConditionChainAnd struct{}

func (ConditionChainAnd) EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error) {
	passedConditions := 0

	for _, condition := range conditions {
		passed, err := condition.Evaluate(input)
		if err != nil {
			return false, err
		}
		if passed {
			passedConditions++
		}
	}

	return passedConditions == len(conditions), nil
}

type ConditionChainOr struct{}

func (ConditionChainOr) EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error) {
	for _, condition := range conditions {
		passed, err := condition.Evaluate(input)
		if err != nil {
			return false, err
		}
		if passed {
			return true, nil
		}
	}

	return false, nil
}
