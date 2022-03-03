package parser

import (
	"gosh/src/lexer"
)

// Parse is created with tokens (by lexer) and creates a executable tree out of it
// by callintg the Parse() function.
type Parser struct {
	tokens []lexer.Token
	pos    int
}

// NewParser creates a new Parser
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens, -1}
}

// Parse parses a list of statements
func (p *Parser) Parse() (*BlockStmt, ParseError) {
	return p.readBlock()
}

// readEvalStmt at pos+1, can return nil
func (p *Parser) readEvalStmt() (EvalStmt, ParseError) {
	var err ParseError
	err = p.tryAdvancePos() //advance onto the character to be scanned
	if err != nil {
		return nil, err
	}

	switch p.tokens[p.pos].TokenType {
	case lexer.TtLBrace:
		return p.readBlock()
	case lexer.TtIf:
		return p.readIfStmt()
	case lexer.TtFor:
		return p.readForStmt()
	default:
		p.pos-- //nothing scanned; decrease pos to stay behind closing braces
		return &ConstStmt{&NilValue{}}, nil
	}
}

func (p *Parser) tryAdvancePos() ParseError {
	if p.pos+1 < len(p.tokens) {
		p.pos++
		return nil
	} else {
		return &UnexpectedEndOfFileError{}
	}
}
