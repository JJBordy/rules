package main

import (
	"encoding/json"
	"fmt"
	"github.com/JJBordy/rules/rules"
	"gopkg.in/yaml.v3"
)

var rulesAsYAML = `
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

	// reading and unmarshalling the rules as rules.RuleInput
	discountRules := make([]rules.RuleInput, 0)
	err := yaml.Unmarshal([]byte(rulesAsYAML), &discountRules)
	if err != nil {
		panic(err)
	}

	// creating a new engine
	engine := rules.NewEngine()
	// creating a set with the rules
	err = engine.CreateSet("discount", discountRules)
	if err != nil {
		panic(err)
	}

	// INPUT A: rule engine input (your custom data)
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

	// evaluating the input against the rules specified for the set
	outputA, err := engine.EvaluateSet("discount", customerA)
	if err != nil {
		panic(err)
	}

	// converting the output to JSON and printing it to the console
	outputJSONforA, err := json.Marshal(outputA)
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT A:", string(outputJSONforA))

	// INPUT B: rule engine input (your custom data)
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

	// evaluating the input against the rules specified for the set
	outputB, err := engine.EvaluateSet("discount", customerB)
	if err != nil {
		panic(err)
	}

	// converting the output to JSON and printing it to the console
	outputJSONforB, err := json.Marshal(outputB)
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT B:", string(outputJSONforB))

	// INPUT C: rule engine input (your custom data)
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

	// evaluating the input against the rules specified for the set
	outputC, err := engine.EvaluateSet("discount", customerC)
	if err != nil {
		panic(err)
	}

	// converting the output to JSON and printing it to the console
	outputJSONforC, err := json.Marshal(outputC)
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT C:", string(outputJSONforC))

	// Output:
	// OUTPUT A: {"packaging":{"style":"simple"}}
	// OUTPUT B: {"discount":{"add":[3]},"packaging":{"style":"premium"}}
	// OUTPUT C: {"discount":{"add":[5,3]},"packaging":{"style":"premium"}}
}
