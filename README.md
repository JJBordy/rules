# simple-rules

Basic minimal rules engine, uses yaml input

It is called 'simple rules' because it will have an input and will produce another; non-recursively.

The reason is to help non-technical people to define conditions based on an input and based on those conditions the
output will be different.

The goal is to have well-defined and well explained schema for the conditions.

### ROADMAP BEFORE 1.0:

* modify the structure for CONDITIONS / CONDITIONS_LIST
* Reference array elements
* add all the built-in functions necessary
* use testify for tests
* add meaningful tests for all .go files
* create good and meaningful examples
* ask feedback before publishing
* publish

### POST 1.0 ROADMAP:

* Rules will be able to have a priority (rules with higher priority will be executed latest and override any previous
  output)
* Rules will gain ID and conditions will be able to reference other rules
* Maybe: full rules engine: will be able to reference the generated output in conditions
* Maybe: output functions (random, time, capitalize string, etc)