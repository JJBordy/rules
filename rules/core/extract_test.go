package core

import (
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "red", extractFieldVal("car.trunk.color", input))
	assert.Equal(t, 31, extractFieldVal("car.roof.resistance", input))
	assert.Equal(t, true, extractFieldVal("car.roof.insured", input))
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
	assert.Equal(t, []any{"Max", "George", "Dory"}, result)
}
