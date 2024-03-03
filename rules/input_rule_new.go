package rules

type RuleInputNew struct {
	Name                string `yaml:"name"`
	ConditionsChainType string `yaml:"chain"`

	Conditions RuleInputConditions `yaml:"conditions"`

	Output    map[string]interface{} `yaml:"output"`
	Map       map[string]interface{} `yaml:"map"`
	OutputMap map[string]string      `yaml:"outputMap"`
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
