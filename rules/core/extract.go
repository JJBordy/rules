package core

import (
	"reflect"
	"strings"
)

// extracts the values from input which contains lists
// car.windows[*].safety.ratings[*].certification will return certifications of all ratings of all windows of the car
func extractFromSlice(path string, input map[string]interface{}) []any {

	pathElems := strings.Split(path, ".")
	resultSlice := make([]any, 0)
	workMap := input

	slicePath := ""

	for pi, pathElem := range pathElems {
		if strings.HasSuffix(pathElem, "[*]") {

			slicePath = strings.TrimSuffix(pathElem, "[*]")

			if arrMap, ok := workMap[slicePath].([]map[string]interface{}); ok {
				for _, arrElem := range arrMap {
					resultSlice = append(resultSlice, extractFromSlice(strings.Join(pathElems[pi+1:], "."), arrElem)...)
				}
			} else if reflect.ValueOf(workMap[slicePath]).Kind() == reflect.Slice {
				sliceVal := reflect.ValueOf(workMap[slicePath])
				for i := 0; i < sliceVal.Len(); i++ {
					resultSlice = append(resultSlice, sliceVal.Index(i).Interface())
				}
			}

		} else if mp, ok := workMap[pathElem].(map[string]interface{}); ok {
			workMap = mp
		} else {
			if singularVal, ok := workMap[pathElem]; ok {
				resultSlice = append(resultSlice, singularVal)
			}
		}
	}

	return resultSlice
}

// extracts the value from input, specified by the nested path, separated by "."
// for example: car.trunk.color
func extractFieldVal(path string, input map[string]interface{}) any {
	workMap := input
	for _, fieldName := range strings.Split(path, ".") {
		if innerMap, ok := workMap[fieldName].(map[string]interface{}); ok {
			workMap = innerMap
		} else {
			return workMap[fieldName]
		}
	}

	return nil
}

func extractFieldList(path string, input map[string]interface{}) []any {
	return extractFromSlice(path, input)
}
