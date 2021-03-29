package parser

import "fmt"

type ParseError interface {
	Error() string
}

type UnsupportedOperationError struct {
	leftType  ValueType
	rightType ValueType
	operator  OperandType
}

func (e *UnsupportedOperationError) Error() string {
	return fmt.Sprintf("Unsupported operation of %s %s %s", e.leftType, e.operator, e.rightType)
}
