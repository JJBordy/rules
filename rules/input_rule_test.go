package rules

import (
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRuleReading(t *testing.T) {

	c, err := os.ReadFile("../testdata/rule_read_test.yaml")
	if err != nil {
		t.Fatal("could not read yaml file")
	}

	var rules []RuleInput

	err = yaml.Unmarshal(c, &rules)
	if err != nil {
		t.Fatal("could not unmarshall yaml content: ", err)
	}

	exampleRule1 := rules[0]
	assert.Equal(t, "Rule Name", exampleRule1.Name)
	assert.Equal(t, "rule-id", exampleRule1.ID)
	assert.Equal(t, 10, exampleRule1.Priority)
	assert.Equal(t, "AND", exampleRule1.ConditionsChain)

	conditions1 := exampleRule1.Conditions[0]
	assert.Equal(t, "customer.balance.usd", conditions1.Input)
	assert.NotNil(t, conditions1.Functions["GREATER"])
	assert.NotNil(t, conditions1.Functions["LESS_THAN"])
	conditions2 := exampleRule1.Conditions[1]
	assert.Equal(t, "customer.name", conditions2.Input)
	assert.NotNil(t, conditions2.Functions["EQUAL"])
	assert.NotNil(t, conditions2.Functions["EQUAL_ANY"])

	exampleRule2 := rules[1]
	assert.Equal(t, "customer business card color", exampleRule2.Name)
	assert.Equal(t, map[string]interface{}{"3": "blue", "5": "red"}, exampleRule2.Map)
	assert.Equal(t, map[string]string{"file.color": "$car.windshield.size"}, exampleRule2.OutputMap)
}
