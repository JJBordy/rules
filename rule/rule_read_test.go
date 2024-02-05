package rule

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRuleReading(t *testing.T) {

	c, err := os.ReadFile("../testdata/example.yaml")
	if err != nil {
		t.Fatal("could not read yaml file")
	}

	var rules []YAML

	err = yaml.Unmarshal(c, &rules)
	if err != nil {
		t.Fatal("could not unmarshall yaml content: ", err)
	}

	fmt.Printf("YAML: %+v\n", rules[0])

	inputVip := map[string]interface{}{
		"customer": map[string]interface{}{
			"balance": map[string]interface{}{
				"usd": "30",
			},
			"name":    "John",
			"surname": "John",
		},
	}

	output := make(map[string]interface{})

	result, err := evaluate(inputVip, rules[0], output)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("result: %+v\n", result)

}

func evaluate(data map[string]interface{}, rule YAML, output map[string]interface{}) ([][2]string, error) {

	var allTrue int

	for _, andRule := range rule.AND {
		ruleElems := strings.Split(andRule, " ")
		input := ruleElems[0]
		comparison := ruleElems[1]
		value := ruleElems[2]

		inputVal := extractFieldVal(input, data)

		if strings.HasPrefix(value, "$") {
			value = extractFieldVal(strings.TrimPrefix(value, "$"), data)
		}

		if comparison == ">" {
			v, err := strconv.ParseFloat(inputVal, 64)
			if err != nil {
				return nil, err
			}
			c, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			if v > c {
				allTrue++
			}
		} else if comparison == "==" {
			if inputVal == value {
				allTrue++
			}
		}
	}

	if rule.ANDMin < allTrue {
		return nil, nil
	}

	finalOutput := make([][2]string, 0)

	for path, value := range rule.Output {
		finalOutput = append(finalOutput, [2]string{path, value})
	}

	return finalOutput, nil
}

func extractFieldVal(path string, input map[string]interface{}) string {
	workMap := input
	for _, fieldName := range strings.Split(path, ".") {
		if val, ok := workMap[fieldName].(map[string]interface{}); ok {
			workMap = val
		} else {
			return fmt.Sprint(workMap[fieldName])
		}
	}

	return ""
}
