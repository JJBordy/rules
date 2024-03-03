package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputFunction(t *testing.T) {

	condition := NewInputFunction("name", []any{"arg"}, func(input any, args []any) (bool, error) {
		return input == "true", nil
	})

	assert.Equal(t, "name", condition.name)
	assert.Equal(t, []any{"arg"}, condition.args)

	result, err := condition.ExecuteFunction("true")
	assert.True(t, result)
	assert.Nil(t, err)

	res, err := condition.ExecuteFunction("false")
	assert.False(t, res)
	assert.Nil(t, err)
}
