package rule

type YAML struct {
	Name     string `yaml:"name"`
	ID       string `yaml:"ID"`
	Priority int    `yaml:"priority"`

	AND    []string `yaml:"AND"`
	OR     []string `yaml:"OR"`
	ANDMin int      `yaml:"AND_MIN"`

	Contains []string `yaml:"CONTAINS"`

	Output map[string]interface{} `yaml:"OUTPUT"`

	MapID string                 `yaml:"MAP-ID"`
	Map   map[string]interface{} `yaml:"MAP"`

	OutputMap []OutputMapYAML `yaml:"OUTPUT-MAP"`

	OutputAppend OutputAppend `yaml:"OUTPUT-APPEND"`
}

type OutputAppendYAML struct {
	Output string              `yaml:"output"`
	Append []map[string]string `yaml:"append"`
}

type OutputMapYAML struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
	MapID  string `yaml:"map"`
}
