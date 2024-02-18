package rules

type RuleInput struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"id"`
	Priority int    `yaml:"priority"`

	ConditionsChain string           `yaml:"chain"`
	Conditions      []Conditions     `yaml:"conditions"`
	ConditionsList  []ConditionsList `yaml:"conditionsList"`

	Output map[string]interface{} `yaml:"output"`

	Map       map[string]interface{} `yaml:"map"`
	OutputMap map[string]string      `yaml:"outputMap"`
}

type Conditions struct {
	Input     string           `yaml:"input"`
	Functions map[string][]any `yaml:"functions"`
}

type ConditionsList struct {
	Inputs    string           `yaml:"inputs"`
	Functions map[string][]any `yaml:"functions"`
}
