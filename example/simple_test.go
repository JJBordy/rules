package example

import (
	"github.com/JJBordy/rules/rules"
	"github.com/go-yaml/yaml"
)

func ExampleEngine_simple() {
	simpleRuleYaml := `
- name: Car color by location
  conditions:
    single:
      - input: car.name
        functions:
          Equal: [ Ferrari ]
      - input: car.height
        functions:
           Greater: 1.2
  OUTPUT:
    paint.trunk: red
    paint.roof: yellow
    tokens: [33, 12]
- name: Car risk by customer properties
  conditions:
    single:
      - input: car.name
        functions:
          Equal: [ Ferrari ]
      - input: car.height
        functions:
           Greater: 1.2
  OUTPUT:
    paint.trunk: red
    paint.roof: yellow
    tokens: [33, 12]
`

	rule := make([]rules.RuleInput, 0)
	err := yaml.Unmarshal([]byte(simpleRuleYaml), &rule)
	if err != nil {
		panic(err)
	}

	// Output: olleh
}
