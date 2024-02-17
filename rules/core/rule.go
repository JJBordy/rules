package core

type Rule struct {
	Name     string
	ID       string
	Priority int

	Conditions []Condition

	ConditionChain *ConditionsChain

	Map map[string]interface{}

	Output    map[string]interface{}
	OutputMap map[string]string
}

func (r Rule) GenerateOutput(input map[string]interface{}) (map[string]interface{}, error) {

	// if OUTPUT_MAP
	// if OUTPUT

	out := make(map[string]interface{})

	rulePasses, err := r.ConditionChain.EvaluateConditions(input, r.Conditions)
	if err != nil {
		return nil, err
	}
	if rulePasses {
		out = r.Output
	}

	return out, nil
}

func (r Rule) DebugOutput(input map[string]interface{}) ([]DebugConditions, map[string]interface{}, error) {
	defer func() {
		r.ConditionChain.TurnDebugOFF()
	}()

	r.ConditionChain.TurnDebugON()
	rulePasses, err := r.ConditionChain.EvaluateConditions(input, r.Conditions)
	if err != nil {
		return nil, nil, err
	}

	if rulePasses {
		return r.ConditionChain.DebugOutput(), r.Output, nil
	} else {
		return r.ConditionChain.DebugOutput(), nil, nil
	}
}
