package rules

import (
	"fmt"
	"github.com/JJBordy/rules/test"
	"github.com/go-yaml/yaml"
	"os"
	"testing"
)

func TestEngine(t *testing.T) {

	engine := NewEngine(EngineConstructorData{})

	fileContent, err := os.ReadFile("../testdata/simple_rule.yaml")
	test.AsserErrtNil(err, t)

	var carRentalRules []RuleInput
	err = yaml.Unmarshal(fileContent, &carRentalRules)
	test.AsserErrtNil(err, t)

	err = engine.CreateSet("car rental", carRentalRules)
	test.AsserErrtNil(err, t)

	carRentalInput := map[string]interface{}{
		"customer": map[string]interface{}{
			"usd": 50,
			"eur": 51,
			"ron": 1000,
		},
	}

	outputOfSet, err := engine.EvaluateSet("car rental", carRentalInput)
	fmt.Printf("setOutput: %+v\n", outputOfSet)

	test.AsserErrtNil(err, t)

	expectedOutput := map[string]interface{}{
		"bonus": map[string]interface{}{
			"points": 5,
			"name":   "super bonus",
		},
	}

	test.AssertEqual(outputOfSet, expectedOutput, t)
}
