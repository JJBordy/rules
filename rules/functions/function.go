package functions

// Function - definition of a function, which takes one element as input
// The concrete elements of this type have the exact implementation of the function
type Function func(input any, args []any) (bool, error)

// FunctionOfList - definition of a function, which takes a list of elements as input
// The concrete elements of this type have the exact implementation of the function
type FunctionOfList func(inputs []any, args []any) (bool, error)
