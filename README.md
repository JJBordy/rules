# simple-rules

Basic minimal rules engine, uses yaml as text input.

It is called 'simple rules' because it will have an input and will produce another; non-recursively.

The reason is to help non-technical people to define conditions based on an input and based on those conditions the
output will be different.

The goal is to have well-defined and well explained schema for the conditions.

## Rule example


### ROADMAP BEFORE 1.0:

* process dates
* for function arguments: allow to pass input values, not just fixed values (Greater: [ config.max ])

### POST 1.0 ROADMAP:

* Rules will be able to have a priority (rules with higher priority will be executed latest and override any previous
  output)
* Rules will gain IDs and conditions will be able to reference other rules (if rule X and rule Y are true/false, then...)
* Possible: full rules engine, will be able to reference the generated output in conditions
* Possible: output functions (random, time, capitalize string, math operations, etc)