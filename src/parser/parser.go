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

func (p *Parser) Parse() (interface{}, ParseError) {
	return nil, nil
}

func (p *Parser) ReadStatement() (ExecutableStatement, ParseError) {
	p.pos++
	switch p.tokens[p.pos].TokenType {
	case lexer.TtPubVarIdent, lexer.TtPrivVarIdent:
		break
	case lexer.TtIdentifier:
		break
	case lexer.TtString:
		break
	case lexer.TtLBrace:
		contents := []ExecutableStatement{}
		for p.tokens[p.pos].TokenType != lexer.TtRBrace {
			statement, err := p.ReadStatement()
			if err != nil {
				return nil, err
			}
			contents = append(contents, statement)
			if p.pos+1 >= len(p.tokens) {
				return nil, &MissingClosingBraceError{}
			}
		}
		return &BlockStatement{contents}, nil
		break
	default:
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, []lexer.TokenType{lexer.TtPubVarIdent, lexer.TtPrivVarIdent, lexer.TtIdentifier, lexer.TtString, lexer.TtLBrace}}
		/*case lexer.TtNumber:
			break
		case lexer.TtLBracket:
			break*/
	}
}
