package core

import (
	"github.com/JJBordy/rules/test"
	"testing"
)

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
