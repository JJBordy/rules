package rules

import (
	"github.com/JJBordy/rules/test"
	"github.com/go-yaml/yaml"
	"os"
	"testing"
)

func TestRuleReading(t *testing.T) {

	c, err := os.ReadFile("testdata/rule_read_test.yaml")
	if err != nil {
		t.Fatal("could not read yaml file")
	}

	var rules []RuleInput

	err = yaml.Unmarshal(c, &rules)
	if err != nil {
		t.Fatal("could not unmarshall yaml content: ", err)
	}

	exampleRule1 := rules[0]
	test.AssertEqual(exampleRule1.Name, "Rule Name", t)
	test.AssertEqual(exampleRule1.ID, "rule-id", t)
	test.AssertEqual(exampleRule1.Priority, 10, t)
	test.AssertEqual(exampleRule1.ConditionsChain, "AND", t)
	test.AssertEqual(exampleRule1.OutputValidation, "validation.output.here", t)

	conditions1 := exampleRule1.Conditions[0]
	test.AssertEqual(conditions1["input"], "customer.balance.usd", t)
	test.AssertEqual(conditions1["GREATER"], []int{10}, t)
	test.AssertEqual(conditions1["LESS_THAN"], []int{100}, t)
	conditions2 := exampleRule1.Conditions[1]
	test.AssertEqual(conditions2["input"], "customer.name", t)
	test.AssertEqual(conditions2["EQUAL"], []string{"$customer.surname"}, t)
	test.AssertEqual(conditions2["EQUAL_ANY"], []string{"Ion", "Vasile", "George"}, t)

	exampleRule2 := rules[1]
	test.AssertEqual(exampleRule2.Name, "customer business card color", t)
	test.AssertEqual(exampleRule2.Map, map[int]string{3: "blue", 5: "red"}, t)
	test.AssertEqual(exampleRule2.OutputMap, map[string]string{"file.color": "$car.windshield.size"}, t)
}

func TestExtractFieldVal(t *testing.T) {
	input := map[string]interface{}{
		"car": map[string]interface{}{
			"trunk": map[string]interface{}{
				"color":  "red",
				"design": "42X",
			},
			"roof": map[string]interface{}{
				"resistance": 31,
				"insured":    true,
			},
		},
	}

	test.AssertEqual(extractFieldVal("car.trunk.color", input), "red", t)
	test.AssertEqual(extractFieldVal("car.roof.resistance", input), 31, t)
	test.AssertEqual(extractFieldVal("car.roof.insured", input), true, t)
}

func TestExtractFromSlice(t *testing.T) {
	input := map[string]interface{}{
		"customer": map[string]interface{}{
			"familyMembers": []map[string]interface{}{
				{
					"name": "Margaret",
					"permit": map[string]interface{}{
						"pets": []map[string]interface{}{
							{
								"name":       "Max",
								"vaccinated": true,
							},
							{
								"name":       "George",
								"vaccinated": false,
							},
						},
					}},
				{
					"name": "Peter",
					"permit": map[string]interface{}{
						"pets": []map[string]interface{}{
							{
								"name":       "Dory",
								"vaccinated": true,
							},
						},
					},
				},
			},
		},
	}

	path := "customer.familyMembers[*].permit.pets[*].name"
	result := extractFromSlice(path, input)
	test.AssertEqual(result, []any{"Max", "George", "Dory"}, t)
}
