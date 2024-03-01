package rules

import (
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRuleReading(t *testing.T) {

	c, err := os.ReadFile("../testdata/rule-better.yml")
	if err != nil {
		t.Fatal("could not read yaml file")
	}

	var rules []RuleInputNew

	err = yaml.Unmarshal(c, &rules)
	if err != nil {
		t.Fatal("could not unmarshall yaml content: ", err)
	}

	assert.Equal(t, "Rule name", rules[0].Name)
	assert.Equal(t, "OR", rules[0].ConditionsChainType)

	// TODO: full input reading test
}
