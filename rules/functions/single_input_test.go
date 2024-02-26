package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultSingleInputFunctions(t *testing.T) {

	defaultFunctions := Default()

	testCases := []struct {
		FuncName     string
		Args         []any
		Input        any
		ExpectOutput bool
		ExpectError  bool
	}{
		{},
	}

	for _, tc := range testCases {
		result, err := defaultFunctions[tc.FuncName](tc.Input, tc.Args)
		if tc.ExpectError {
			assert.Equal(t, tc.ExpectOutput, result)
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}
