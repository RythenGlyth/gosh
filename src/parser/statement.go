package parser

import "io"

// EvalStmt is an evaluable statement.
type EvalStmt interface {
	// Value is not nil as long as RuntimeError is nil
	Eval() (Value, RuntimeError)

	Debug(out io.Writer, indent int, symbol string)
}
