package test

import (
	"fmt"
	"runtime"
	"testing"
)

func AssertEqual(got, expected any, t *testing.T) {
	if fmt.Sprint(got) != fmt.Sprint(expected) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("Expected %v, got %v; %s:%d", expected, got, file, line)
	}
}
