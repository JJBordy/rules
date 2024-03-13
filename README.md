# rules

[![Go Reference](https://pkg.go.dev/badge/github.com/JJBordy/rules.svg)](https://pkg.go.dev/github.com/JJBordy/rules)

Minimal rules engine.

It's purpose is to allow non-programmers to edit and provide business rules which will generate a specific output.

It uses yaml because of a fixed syntax which people can refer to online (not a scripting language). It still has a lot
of verbosity to facilitate development of a GUI/client which can read and edit them. Also that way they are easy to
store/transfer.

Check the example in [example_test.go](example_test.go)

## short description

Could also be considered as a 'conditional output' engine, as the rules receive an input (`map[string]any`) and
generates an output if the rule is fulfilled. The output is different from the input.

The rules are combined into a set, thus passing an input to a set of rules returns the combined output of all the rules
that passed for the input.

Each rule can have multiple conditions. If they are fulfilled then the output of the rule will be appended/written.

Has a debug capability. Alongside the output, the engine will return a set of all the rules/conditions which didn't pass
for a given input.

## what it is not

The input is separated from the output, only the input is evaluated.

There is no relation between the rules, no referring of one rule to another, no hierarchy of rules.

It does not perform any calculations, except 'SUM' from a list of values.

## rule anatomy

Below will go through a rule line by line and will describe each of its elements. The rule can be found
in [testdata/rule-input.yaml](testdata/rule-input.yaml). Some values may not be exactly as there in order to showcase
additional functionality.

#### *name*

`name: Rule name`

The name of the rule. Should be unique if you want to use the debug capability (retrieve the rules/conditions which
failed)

#### *chain*

`chain: OR`

Each rule has a set of conditions. By default, all conditions have to be fulfilled in order for the rule to generate its
output (the default `chain` value is `AND`).

The `chain` property can have any property of the logic gates operators. Documentation: [rules/core/condition_chain.go](rules/core/condition_chain.go)

#### *conditions*

`conditions:`

The conditions of the rule. They can be of 3 types: single, list and aggregate. Will examine each of the types below.

#### *single conditions*

```
  conditions:
    single:
      - input: customer.name
        functions:
          Equal: [ John ]
```

The `single` property has a list of single conditions. Each single condition has an input path (fields separated by a
dot) which leads to a value in the input to the engine.

The `functions` property contains a map of functions which evaluate the value specified in the input path. Multiple
functions can be added under the same conditions.
The list of all functions are here: [rules/functions/single_input.go](rules/functions/single_input.go). You can also create your own functions and add
them to the engine, with: `rules.NewEngineCustom()`

The arguments for all the functions are specified inside square brackets. Even if there are no arguments, the square
brackets should be there.

#### *list conditions*

```
  list:
    - inputs: some.list.is.here[*].name
      constraints:
        AtLeast: [ 3 ]
      functions:
        Between: [ 10, 100 ]
```

The `list` property specifies list conditions, those which refer not to single field values but to lists of values.

The `inputs` property specifies the path to the list of values in the input. The field which contains a list is
specified with `[*]`.

Each list condition has `constraints` and `functions` properties. The `functions` are the same map of functions which
are in the single conditions.

By default, all elements of the list will have to pass all the functions. If you want to change that logic and have only
a certain number of elements in the list pass the functions, you can use `constraints`.
All of them are documented here: [rules/functions/list_constraints.go](rules/functions/list_constraints.go)

#### *aggregate conditions*

```
  aggregate:
    - inputs: some.list.is.here[*].rating
      type: MIN
      functions:
        Greater: [ 100 ]
```

The `aggregate` property specifies aggregate conditions, they also refer to lists of values. While the `list` conditions
refer to each element one by one, the `aggregate` conditions refer to them as a whole.

The `type` property specifies the aggregate type. All possible values are documented
here: [rules/functions/aggregate.go](rules/functions/aggregate.go)

The `functions` property specifies the functions that are used to evaluate the result of the aggregate. They are the
same as for single and list conditions.

#### *OUTPUT*

```
  OUTPUT:
    file.size: 31
    file.name: some file name
    any.other.key: value
    some.list.of.values: [ a, b]
```

Under the `OUTPUT` property you can specify a map of values which will be returned in case the rule succeeds.

When evaluating a set of rules, the engine combines the output of all the rules which succeeded into a single map and
returns it.

The output has the fields nesting specified with `.` (like json path).

You can also output lists of simple values (no structs/maps). You just specify the value in square brackets.

In case that multiple rules change the same field - the latest rule specified will override the previous ones.

In case of lists - the values are appended into a single list. This way, if you need each rule to increment/decrement a
value, you can use the values from an output list.

#### *MAP* and *OUTPUT_MAP*

```
  MAP:
    new: white
    regular: blue
    vip: golden
  OUTPUT_MAP:
    card.color: $customer.role
    group.roles.given: [ $customer.role ]
```

The `MAP` property specifies a map of key/value pairs. In the `OUTPUT_MAP` property you can additionaly specify the
output of the rule (`OUTPUT` and `OUTPUT_MAP` are complementary, you can have only one of them or both).
`OUTPUT_MAP` requires `MAP` to be present. You specify the output as in `OUTPUT`, only instead of values you specify a
path inside the input.
Thus, the value from the input corresponding to the map key will be written in the output.

## roadmap:

* process dates
* for function arguments: allow to pass input values, not just fixed values (ex function: `Greater: [ $customer.age ]`)
* rules will be able to have a priority (rules with higher priority will be executed latest and override any previous
  output)
* Rules will gain IDs and conditions will be able to reference other rules (if rule X and rule Y are true/false,
  then...)
* Possible: full rules engine, will be able to reference the generated output in conditions
* Possible: output functions (random, time, capitalize string, math operations, etc)
