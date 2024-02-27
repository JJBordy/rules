package functions

// constraints examples: at least fraction, at most fraction, all, none

const (
	// All - all elements of the list should pass the functions (default)
	All = "All"
	// AtLeast - minimum elements in list to pass the functions
	AtLeast = "AtLeast"
	// AtMost - maximum elements in list to pass the functions
	AtMost = "AtMost"
	// Exactly - exact number of elements in list to pass the functions
	Exactly = "Exactly"
	// None - no elements in list should pass the functions
	None = "None"
	// AtLeastFraction - minimum fraction of elements in list to pass the functions (1,2) = 50%, (3,4) - 75%
	AtLeastFraction = "AtLeastFraction"
	// AtMostFraction - maximum fraction of elements in list to pass the functions (1,2) = 50%, (3,4) - 75%
	AtMostFraction = "AtMostFraction"
)

type ListFunctionConstraint func(listTotal, passedTotal int, args []int) bool

func AllListFunctionConstraints() map[string]ListFunctionConstraint {
	allConstraints := make(map[string]ListFunctionConstraint)

	allConstraints[All] = func(listTotal, passedTotal int, args []int) bool {
		return passedTotal == listTotal
	}

	allConstraints[AtLeast] = func(listTotal, passedTotal int, args []int) bool {
		return args[0] >= passedTotal
	}

	allConstraints[AtMost] = func(listTotal, passedTotal int, args []int) bool {
		return args[0] <= passedTotal
	}

	allConstraints[Exactly] = func(listTotal, passedTotal int, args []int) bool {
		return args[0] == passedTotal
	}

	allConstraints[None] = func(listTotal, passedTotal int, args []int) bool {
		return passedTotal == 0
	}

	allConstraints[AtLeastFraction] = func(listTotal, passedTotal int, args []int) bool {
		expected := float64(listTotal/args[1]) * float64(args[0])
		return float64(passedTotal) >= expected
	}

	allConstraints[AtMostFraction] = func(listTotal, passedTotal int, args []int) bool {
		expected := float64(listTotal/args[1]) * float64(args[0])
		return float64(passedTotal) <= expected
	}

	return allConstraints
}
