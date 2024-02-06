package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func BuildOutput(outputResults []map[string]any) (output map[string]any, err error) {
	defer func() {
		if pan := recover(); pan != nil {
			err = errors.New(fmt.Sprint(pan))
		}
	}()

	output = make(map[string]any)

	for _, result := range outputResults {
		for k, v := range result {
			output = setValue(output, k, v)
		}
	}

	return output, err
}

func setValue(m map[string]any, path string, val any) map[string]any {
	if m == nil {
		m = make(map[string]any)
	}
	if strings.Contains(path, ".") {
		splitPath := strings.Split(path, ".")
		var mapToPass map[string]any
		if mapPart, ok := m[splitPath[0]].(map[string]any); ok {
			mapToPass = mapPart
		}
		m[splitPath[0]] = setValue(mapToPass, strings.Join(splitPath[1:], "."), val)
	} else {
		switch reflect.ValueOf(val).Kind() {
		case reflect.Slice:
			if m[path] == nil {
				m[path] = val
			} else {
				for i := 0; i < reflect.ValueOf(val).Len(); i++ {
					slice := reflect.Append(reflect.ValueOf(m[path]), reflect.ValueOf(val).Index(i))
					reflect.ValueOf(m).SetMapIndex(reflect.ValueOf(path), slice)
				}
			}
		default:
			m[path] = val
		}
	}
	return m
}
