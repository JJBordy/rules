# simple-rules

Basic minimal rules engine, uses yaml input

It is called 'simple rules' because it will have an input and will produce another; non-recursively.

The reason is to help non-technical people to define conditions based on an input and based on those conditions the
output will be different.

The goal is to have well-defined and well explained schema for the conditions.

In the future should provide:
* validation of rules with detailed explanation of why it failed.
* validation of input provided in case it is not compliant with a rule.
* able to tell which rules failed and why, which rules succeeded and why.