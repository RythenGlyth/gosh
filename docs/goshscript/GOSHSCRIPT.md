# GOSHScript

A script consists of evaluable statements.

If seperations can't be known by context, you need to use semicolons as seperators. New lines also act like semicolons.

An evaluable statements `<ES>` can be
1. Block statement
1. If statement
1. For statement (`while` is a synonym for `for`)
1. Assignment statement
1. Function Call
1. Command Call
1. Calculated Statement (`<ES> <OP> <ES>`, where `<OP>` is any operator)
1. Variable
1. Constant

## Block Statements

A block statement is a group of executable statements are grouped using braces `{ES ES}`.
Block statements are executable statements and evaluated to nothing.

## If Statements

If statements must have a conditional statement and a true statement, and can have a false statement. 

The conditional statement can be of any value.
If the value is a ConditionStatement, it will be evaluated to boolean.
Everything else will be evaluated to false if 0, else to true.

The true and false statements are a single statement (e.g. block statement).

 - `if (VS) ES`
 - `if (VS) ES else ES`

## For Statements

For statements are inspired by Go:
 - `for (ES; VS; ES) ES`
 - `for (ES) ES`

They must also have a body.

## Assignment

Assigning a variable to a value statement, e.g. `VAR = VS`

## Value Statements

 - Constant Value
 - (Value Statement) <OP> (Value Statement) - number of VS depends on the operator
 - Variable
 - Executable Statements
