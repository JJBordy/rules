package rules

type RuleInput struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id"`
	Priority int    `yaml:"priority"`

	ConditionsChain   string                   `yaml:"COND_CHAIN"`
	ConditionsMinimum int                      `yaml:"COND_MIN"`
	Conditions        []map[string]interface{} `yaml:"COND"`

	Output           map[string]interface{} `yaml:"OUTPUT"`
	OutputValidation string                 `yaml:"OUTPUT_VALIDATION"`

	Map       map[string]interface{} `yaml:"MAP"`
	OutputMap map[string]string      `yaml:"OUTPUT_MAP"`
}

type ConditionFunc func(input any, args []any) (bool, error)

type Rule struct {
	Name     string
	ID       string
	Priority int

	conditions []Condition

	conditionChain ConditionChain

	Map map[string]interface{}

	Output    map[string]interface{}
	OutputMap map[string]interface{}
}

func (r Rule) GenerateOutput(input map[string]interface{}) (map[string]interface{}, error) {
	rulePasses, err := r.conditionChain.EvaluateConditions(input, r.conditions)
	if err != nil {
		return nil, err
	}
	if rulePasses {
		// if outputMap present, merge it and return it too
		// also, output from validation
		return r.Output, nil
	}
	return nil, nil
}

// ConditionChain - AND, OR, XAND, NOR, etc
type ConditionChain interface { // may have condition chain type - as interface
	EvaluateConditions(input map[string]interface{}, conditions []Condition) (bool, error)
}

type Condition struct {
	inputValue     any
	conditionFuncs []ConditionFunc
}

func ParseRuleInput(ruleInput RuleInput) (Rule, error) {
	return Rule{}, nil
}
