package parser

import (
	"fmt"
	"gosh/src/lexer"
)

type RuntimeError interface {
	error
}

type ParseError interface {
	error
}

type UnexpectedTokenError struct {
	is       lexer.TokenType
	expected string
}

func (e *UnexpectedTokenError) Error() string {
	return fmt.Sprintf("Unexpected token: %s. expected: %s", e.is, e.expected)
}

type UnexpectedEndOfFileError struct {
}

func (e *UnexpectedEndOfFileError) Error() string {
	return fmt.Sprintf("Unexpected end of file")
}
