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
	OutputMap map[string]interface{} `yaml:"OUTPUT_MAP"`
}
