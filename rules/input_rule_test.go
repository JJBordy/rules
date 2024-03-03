package rules

import (
	"github.com/JJBordy/rules/rules/functions"
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRuleReading(t *testing.T) {

	c, err := os.ReadFile("../testdata/rule-input.yaml")
	if err != nil {
		t.Fatal("could not read yaml file")
	}

	var rules []RuleInput

	err = yaml.Unmarshal(c, &rules)
	if err != nil {
		t.Fatal("could not unmarshall yaml content: ", err)
	}

	// rule
	assert.Equal(t, 1, len(rules))
	rule := rules[0]
	assert.Equal(t, "Rule name", rule.Name)
	assert.Equal(t, "OR", rule.ConditionsChainType)

	// single
	assert.Equal(t, 1, len(rule.Conditions.SingleInputConditions))
	siCondition := rule.Conditions.SingleInputConditions[0]
	assert.Equal(t, "customer.name", siCondition.InputPath)
	assert.Equal(t, 1, len(siCondition.Functions))
	assert.Equal(t, []any{"John"}, siCondition.Functions[functions.Equal])

	// list
	assert.Equal(t, 1, len(rules[0].Conditions.ListInputConditions))
	listCondition := rules[0].Conditions.ListInputConditions[0]
	assert.Equal(t, "some.list.is.here[*].name", listCondition.InputsPath)
	assert.Equal(t, 1, len(listCondition.Constraints))
	assert.Equal(t, []any{3}, listCondition.Constraints[functions.AtLeast])
	assert.Equal(t, 1, len(listCondition.Functions))
	assert.Equal(t, []any{10, 100}, listCondition.Functions[functions.Between])

	// aggregate
	assert.Equal(t, 1, len(rules[0].Conditions.ListAggregateConditions))
	aggregateCondition := rules[0].Conditions.ListAggregateConditions[0]
	assert.Equal(t, "some.list.is.here[*].rating", aggregateCondition.InputsPath)
	assert.Equal(t, functions.MIN, aggregateCondition.AggregateType)
	assert.Equal(t, 1, len(aggregateCondition.Functions))
	assert.Equal(t, []any{100}, aggregateCondition.Functions[functions.Greater])
}
