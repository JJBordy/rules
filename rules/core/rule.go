package core

import (
	"fmt"
	"strings"
)

type RuleNew struct {
	Name string

	ConditionChain *ConditionsChain

	Conditions []Condition

	Map map[string]interface{}

	Output    map[string]interface{}
	OutputMap map[string]string
}

func (r RuleNew) GenerateOutput(input map[string]interface{}) (map[string]interface{}, error) {
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

	// needs to be separated
	for outPath, inPath := range r.OutputMap {
		if strings.Contains(inPath, "[*]") {
			fieldListValues := extractFromSlice(inPath, input)
			if fieldListValues == nil {
				continue
			}
			resultList := make([]interface{}, 0)
			for _, mapKey := range fieldListValues {
				if r.Map[fmt.Sprint(mapKey)] == nil {
					continue
				}
				resultList = append(resultList, r.Map[fmt.Sprint(mapKey)])
			}
			out[outPath] = resultList

		} else {
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
	}

	return out, nil
}

func (r RuleNew) DebugOutput(input map[string]interface{}) ([]DebugCondition, map[string]interface{}, error) {
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
