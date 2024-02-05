package rule

type Rule struct {
	Name           string
	ID             string
	ConditionChain string
	Conditions     []Condition
}

type Condition struct {
	InputPath string
	Func      Function
	ListFunc  FunctionList
	Values    []string
}

type Function interface {
	Apply(value interface{}) bool
}

type FunctionList interface {
	Apply(values []interface{}) bool
}

type OutputSimple struct {
}

type OutputAppend struct {
}

type OutputMap struct {
}
