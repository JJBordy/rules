package rules

type RuleInput struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id"`
	Priority int    `yaml:"priority"`

	ConditionsChain string                   `yaml:"COND_CHAIN"`
	Conditions      []map[string]interface{} `yaml:"COND"`

	Output map[string]interface{} `yaml:"OUTPUT"`

	Map       map[string]interface{} `yaml:"MAP"`
	OutputMap map[string]string      `yaml:"OUTPUT_MAP"`
}
