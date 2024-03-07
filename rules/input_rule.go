package rules

type RuleInput struct {
	Name                string `yaml:"name"`
	ConditionsChainType string `yaml:"chain"`

	Conditions RuleInputConditions `yaml:"conditions"`

	Output    map[string]interface{} `yaml:"OUTPUT"`
	Map       map[string]interface{} `yaml:"MAP"`
	OutputMap map[string]string      `yaml:"OUTPUT_MAP"`
}

type RuleInputConditions struct {
	SingleInputConditions   []ConditionSingleInput   `yaml:"single"`
	ListInputConditions     []ConditionListInput     `yaml:"list"`
	ListAggregateConditions []ConditionListAggregate `yaml:"aggregate"`
}

type ConditionSingleInput struct {
	InputPath string           `yaml:"input"`
	Functions map[string][]any `yaml:"functions"`
}

type ConditionListInput struct {
	InputsPath  string           `yaml:"inputs"`
	Functions   map[string][]any `yaml:"functions"`
	Constraints map[string][]any `yaml:"constraints"`
}

type ConditionListAggregate struct {
	InputsPath    string           `yaml:"inputs"`
	AggregateType string           `yaml:"type"`
	Functions     map[string][]any `yaml:"functions"`
}
