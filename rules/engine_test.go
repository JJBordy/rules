package rules

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEngine(t *testing.T) {

	engine := NewEngine(EngineConstructorData{})

	fileContent, err := os.ReadFile("../testdata/simple_rule.yaml")
	assert.Nil(t, err)

	var carRentalRules []RuleInput
	err = yaml.Unmarshal(fileContent, &carRentalRules)
	assert.Nil(t, err)

	err = engine.CreateSet("car rental", carRentalRules)
	assert.Nil(t, err)

	// example input
	carRentalInput := map[string]interface{}{
		"customer": map[string]interface{}{
			"usd":     50,
			"eur":     51,
			"ron":     1000,
			"age":     70,
			"name":    "George",
			"friends": []string{"Vasile", "Ion", "Michael"},
		},
	}

	outputOfSet, err := engine.EvaluateSet("car rental", carRentalInput)
	fmt.Printf("setOutput: %+v\n", outputOfSet)
	assert.Nil(t, err)

	expectedOutput := map[string]interface{}{
		"bonus": map[string]interface{}{
			"points":    5,
			"name":      "super bonus",
			"superName": "Super George",
			"friends":   []interface{}{"Super Vasile", "Super Ion", "Super Mike"},
		},
	}
	assert.Equal(t, expectedOutput, outputOfSet)
}
