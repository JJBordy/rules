package core

import (
	"github.com/JJBordy/rules/rules/functions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleInputCondition(t *testing.T) {
	funcs := functions.Default()

	sic := NewSingleInputCondition("key", []InputFunction{NewInputFunction("funcName", []any{"eqVal"}, funcs[functions.Equal])})

	res, err := sic.Evaluate(map[string]any{"key": "eqVal"})
	assert.Nil(t, err)
	assert.True(t, res)

	deb := sic.DebugInfo()
	assert.Equal(t, "key", deb.InputPath)
	assert.Equal(t, "funcName", deb.Functions[0].Name)
	assert.Equal(t, 1, len(deb.Functions[0].Args))
	assert.Equal(t, 1, len(deb.Functions))
	assert.Equal(t, "eqVal", deb.Functions[0].Args[0])

}
