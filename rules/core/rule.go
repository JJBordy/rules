package core

import "fmt"

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
	out := make(map[string]interface{})

	rulePasses, err := r.ConditionChain.EvaluateConditions(input, r.Conditions)
	if err != nil {
		return nil, err
	}
	if rulePasses {
		if r.Output != nil {
			out = r.Output
		}
	}

	if r.OutputMap == nil || r.Map == nil {
		return out, nil
	}

	for outPath, inPath := range r.OutputMap {
		fieldValue := extractFieldVal(inPath, input)
		if fieldValue == nil {
			continue
		}
		mapKey := fmt.Sprint(fieldValue)
		if r.Map[mapKey] == nil {
			continue
		}
		out[outPath] = r.Map[mapKey]
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
