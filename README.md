# rules

Minimal rules engine, uses yaml as text input.
Could also be considered as a 'conditional output' engine, as the rules receive an input and append to an output if the
rule is fulfilled.

Check the example file to see how it works.

## Rule anatomy

Below will go through a rule line by line and will describe each of its elements.

## Roadmap before 1.0:

* process dates
* for function arguments: allow to pass input values, not just fixed values (Greater: [ config.max ])

## Roadmap post 1.0:

* Rules will be able to have a priority (rules with higher priority will be executed latest and override any previous
  output)
* Rules will gain IDs and conditions will be able to reference other rules (if rule X and rule Y are true/false,
  then...)
* Possible: full rules engine, will be able to reference the generated output in conditions
* Possible: output functions (random, time, capitalize string, math operations, etc)