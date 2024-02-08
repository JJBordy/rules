package rules

type Engine struct {
	conditionChains map[string]ConditionChain
	ruleSets        map[string][]Rule
	conditions      map[string]ConditionFunc
}

func NewEngine() *Engine {
	return &Engine{
		ruleSets:   make(map[string][]Rule),
		conditions: make(map[string]ConditionFunc),
	}
}

func (e *Engine) AddCondition(conditionName string, condition ConditionFunc) {
	e.conditions[conditionName] = condition
}

func (e *Engine) AddToSet(setName string, rules []RuleInput) {

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
