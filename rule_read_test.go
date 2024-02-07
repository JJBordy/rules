package rules

import (
	"github.com/JJBordy/rules.git/test"
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
	test.AssertEqual(exampleRule1.ConditionsMinimum, 2, t)
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

//func TestAppend(t *testing.T) {
//	yamA := `
//some:
//  stuff: true
//  things:
//    look: nice
//    are: [good]
//`
//
//	yamB := `
//some:
//  otherStuff: false
//  things:
//    seem: cool
//    are: [bad]
//`
//
//	var result map[string]interface{}
//
//	err := yaml.Unmarshal([]byte(yamA), &result)
//	if err != nil {
//		t.Fatal(err)
//	}
//	err = yaml.Unmarshal([]byte(yamB), &result)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Printf("%+v\n", result)
//}

// TODO: it works, should be adjusted for actual rule model
//func evaluate(data map[string]interface{}, rule RuleInput, output map[string]interface{}) (map[string]interface{}, error) {
//
//	var allTrue int
//
//	for _, andRule := range rule.Conditions {
//		ruleElems := strings.Split(andRule., " ")
//		input := ruleElems[0]
//		comparison := ruleElems[1]
//		value := ruleElems[2]
//
//		inputVal := extractFieldVal(input, data)
//
//		if strings.HasPrefix(value, "$") {
//			value = extractFieldVal(strings.TrimPrefix(value, "$"), data)
//		}
//
//		if comparison == ">" {
//			v, err := strconv.ParseFloat(inputVal, 64)
//			if err != nil {
//				return nil, err
//			}
//			c, err := strconv.ParseFloat(value, 64)
//			if err != nil {
//				return nil, err
//			}
//			if v > c {
//				allTrue++
//			}
//		} else if comparison == "==" {
//			if inputVal == value {
//				allTrue++
//			}
//		}
//	}
//
//	if rule.ANDMin < allTrue {
//		return nil, nil
//	}
//
//	//finalOutput := make([][2]string, 0)
//
//	//for path, value := range rule.Output {
//	//	finalOutput = append(finalOutput, [2]string{path, value})
//	//}
//
//	return output, nil
//}
//
//func extractFieldVal(path string, input map[string]interface{}) string {
//	workMap := input
//	for _, fieldName := range strings.Split(path, ".") {
//		if val, ok := workMap[fieldName].(map[string]interface{}); ok {
//			workMap = val
//		} else {
//			return fmt.Sprint(workMap[fieldName])
//		}
//	}
//
//	return ""
//}
