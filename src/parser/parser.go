package parser

import (
	"gosh/src/lexer"
)

// Parse is created with tokens (by lexer) and creates a executable tree out of it
// by callintg the Parse() function.
type Parser struct {
	tokens *[]lexer.Token
}

// NewParser creates a new Parser
func NewParser(tokens *[]lexer.Token) *Parser {
	return &Parser{tokens}
}

func (p *Parser) Parse() (interface{}, ParseError) {
	return nil, nil
}
