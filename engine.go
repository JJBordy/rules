package rules

type Engine struct {
	conditionChains        map[string]ConditionChain
	ruleSets               map[string][]Rule
	conditionFunctions     map[string]ConditionFunc
	conditionListFunctions map[string]ConditionFuncOfList
}

func NewEngine(userFunctions map[string]ConditionFunc, userListFunctions map[string]ConditionFuncOfList) *Engine {
	e := Engine{
		ruleSets: make(map[string][]Rule),
	}

	functions := defaultFunctions()
	for k, v := range userFunctions {
		functions[k] = v
	}

	listFunctions := defaultListFunctions()
	for k, v := range userListFunctions {
		listFunctions[k] = v
	}

	e.conditionFunctions = functions
	e.conditionListFunctions = listFunctions

	return &e
}

func (e *Engine) CreateSet(setName string, ruleInputs []RuleInput) error {
	parsedRules := make([]Rule, 0)

	for _, r := range ruleInputs {
		parsedRule, err := ParseRuleInput(r)
		if err != nil {
			return err
		}
		parsedRules = append(parsedRules, parsedRule)
	}

	e.ruleSets[setName] = parsedRules

	return nil
}

func (e *Engine) EvaluateSet(setName string, input map[string]interface{}) (map[string]interface{}, error) {

	outputs := make([]map[string]interface{}, 0)

	for _, rule := range e.ruleSets[setName] {
		output, err := rule.GenerateOutput(input)
		if err != nil {
			return nil, err
		}
		if output != nil {
			outputs = append(outputs, output)
		}
	}

	return BuildOutput(outputs)
}

func defaultFunctions() map[string]ConditionFunc {
	return make(map[string]ConditionFunc)
}

func defaultListFunctions() map[string]ConditionFuncOfList {
	return make(map[string]ConditionFuncOfList)
}

func defaultConditions() map[string]Condition {
	return make(map[string]Condition)
}
