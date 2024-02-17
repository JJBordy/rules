### Default functions for single input

#### General functions

* `EMPTY:` `[any]` - the input is empty
* `NONEMPTY:` `[any]` - the input is not empty
* `EQUAL:` `[any]` - the input is equal to the argument

#### Numeric functions

* `GREATER:` `[nr]` - the input is a number and is greater than the argument
* `GREATER_EQ:` `[nr]` - the input is a number and is greater or equal to the argument
* `LOWER:` `[nr]` - the input is a number and is lower than the argument
* `LOWER_EQ:` `[nr]` - the input is a number and is lower or equal to the argument
* `BETWEEN:` `[nr1, nr2]` - the input is a number and is between the two arguments
* `BETWEEN_EQ:` `[nr1, nr2]` - the input is a number and is between or equal to the two arguments
* `NOT_BETWEEN:` `[nr1, nr2]` - the input is a number and is not between the two arguments
* `NOT_BETWEEN_EQ:` `[nr1, nr2]` - the input is a number and is not between or equal to the two arguments

#### String functions

* `EQUAL_IGNORE_CASE:` `[string]` - the input is equal to the argument case-insensitively
* `EQUAL_ANY:` `[string1, string2....]` - the input is equal to any of the arguments
* `NOT_EQUAL_ANY:` `[string1, string2....]` - the input is not equal to any of the arguments
* `STARTS_WITH:` `[string]` - the input starts with the argument
* `STARTS_WITH_IGNORE_CASE:` `[string]` - the input starts with the argument case-insensitively
* `ENDS_WITH:` `[string]` - the input ends with the argument
* `ENDS_WITH_IGNORE_CASE:` `[string]` - the input ends with the argument case-insensitively
* `CONTAINS:` `[string]` - the input contains the argument
* `CONTAINS_IGNORE_CASE:` `[string]` - the input contains the argument case-insensitively