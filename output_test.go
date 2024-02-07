package rules

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"strings"
	"testing"
)

func TestBuildOutput(t *testing.T) {
	// for overriding results: the results will be sorted in order of their priority, so it will happen automatically
	results := []map[string]any{
		{"file.priority": "low"},
		{"file.priority": "high"},
		{"nested.value.here": "A"},
		{"nested.slice.here": []string{"A", "B", "C"}},
		{"nested.slice.here": []string{"D", "E"}},
		{"nested.slice.here": []string{"F", "G"}},
		{"numbers.here": []int{11, 22}},
		{"numbers.here": []int{33, 44}},
	}

	expectedYamlOutput := `file:
  priority: high
nested:
  slice:
    here:
    - A
    - B
    - C
    - D
    - E
    - F
    - G
  value:
    here: A
numbers:
  here:
  - 11
  - 22
  - 33
  - 44`

	output, err := BuildOutput(results)
	if err != nil {
		t.Fatal(err, "expected no error when building output")
	}

	yamlOutput, err := yaml.Marshal(output)
	if err != nil {
		t.Fatal(err, "expected no error when marshaling output into YAML")
	}

	strYaml := string(yamlOutput)

	if strings.EqualFold(expectedYamlOutput, strYaml) {
		fmt.Println(string(yamlOutput))
		t.Fatal("expected yaml output to be as the defined one")
	}
}
