package rules

import (
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
)

var simpleRuleYaml = `
- name: Setting default values for the output
  conditions:
    single:
      - input: customer.name
        functions:
          NonEmpty: []
  OUTPUT:
    packaging.style: simple
- name: Vulnerable people packaging
  chain: OR
  conditions:
    single:
      - input: customer.age
        functions:
          GreaterEq: [ 65 ]
      - input: customer.disability
        functions:
          Equal: [ true ]
  OUTPUT:
    packaging.style: easy
- name: Discount based on money spent on all orders
  conditions:
    aggregate:
      - inputs: customer.orders[*].price
        type: SUM
        functions:
          Greater: [ 8000 ]
  OUTPUT:
    discount.add: [ 5 ]
    packaging.style: premium
- name: Discount based on large orders
  conditions:
    list:
      - inputs: customer.orders[*].price
        constraints:
          AtLeast: [ 2 ]
        functions:
          Greater: [ 1000 ]
  OUTPUT:
    discount.add: [ 3 ]
    packaging.style: premium
`

func ExampleEngine_simple() {

	discountRules := make([]RuleInput, 0)
	err := yaml.Unmarshal([]byte(simpleRuleYaml), &discountRules)
	if err != nil {
		panic(err)
	}

	engine := NewEngine()
	err = engine.CreateSet("discount", discountRules)
	if err != nil {
		panic(err)
	}

	customerA := map[string]any{
		"customer": map[string]any{
			"name":       "John",
			"age":        30,
			"disability": false,
			"orders": []map[string]any{
				{
					"price": 500,
				},
			},
		},
	}

	outputA, err := engine.EvaluateSet("discount", customerA)
	if err != nil {
		panic(err)
	}

	outputJSONforA, err := json.Marshal(outputA)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprint(string(outputJSONforA)))

	customerB := map[string]any{
		"customer": map[string]any{
			"name":       "Mike",
			"age":        65,
			"disability": false,
			"orders": []map[string]any{
				{
					"price": 2000,
				},
				{
					"price": 2300,
				},
			},
		},
	}

	outputB, err := engine.EvaluateSet("discount", customerB)
	if err != nil {
		panic(err)
	}

	outputJSONforB, err := json.Marshal(outputB)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprint(string(outputJSONforB)))

	customerC := map[string]any{
		"customer": map[string]any{
			"name":       "Stan",
			"age":        24,
			"disability": true,
			"orders": []map[string]any{
				{
					"price": 5000,
				},
				{
					"price": 5000,
				},
				{
					"price": 5000,
				},
			},
		},
	}

	outputC, err := engine.EvaluateSet("discount", customerC)
	if err != nil {
		panic(err)
	}

	outputJSONforC, err := json.Marshal(outputC)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprint(string(outputJSONforC)))

	// Output: {"packaging":{"style":"simple"}}
	// {"discount":{"add":[3]},"packaging":{"style":"premium"}}
	// {"discount":{"add":[5,3]},"packaging":{"style":"premium"}}
}
