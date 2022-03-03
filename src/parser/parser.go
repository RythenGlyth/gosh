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
	case lexer.TtIf:

		break
	case lexer.TtFor:
		break
	default:
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, []lexer.TokenType{lexer.TtPubVarIdent, lexer.TtPrivVarIdent, lexer.TtIdentifier, lexer.TtString, lexer.TtLBrace}}
		/*case lexer.TtNumber:
			break
		case lexer.TtLBracket:
			break*/
	}
}

func (p *Parser) readValueStatement() (ValueStatement, ParseError) {
	/*switch p.tokens[p.pos].TokenType {
	case lexer.TtTrue:
		return &ConstantConditionStatement{true}, nil
	case lexer.TtFalse:
		return &ConstantConditionStatement{false}, nil

	//Rest Depending on value reading
	}*/
valueStatementReaderForLoop:
	for p.pos+1 < len(p.tokens) {
		p.pos++
		switch p.tokens[p.pos].TokenType {
		case lexer.TtLParen:
			break
		case lexer.TtString:
			break
		case lexer.TtNumber:
			break
		case lexer.TtTrue:
			break
		case lexer.TtFalse:
			break
		case lexer.TtLBrace:
			break
		case lexer.TtLBracket:
			break
		case lexer.TtQuestion:
			break
		case lexer.TtColon:
			break
		case lexer.TtPercent:
			break
		case lexer.TtStar:
			break
		case lexer.TtSlash:
			break
		case lexer.TtPlus:
			break
		case lexer.TtMinus:
			break
		case lexer.TtPlusPlus:
			break
		case lexer.TtMinusMinus:
			break
		case lexer.TtAndAnd:
			break
		case lexer.TtBarbar:
			break
		case lexer.TtEqEq:
			break
		case lexer.TtBangEq:
			break
		case lexer.TtLtEq:
			break
		case lexer.TtGtEq:
			break
		default:
			break valueStatementReaderForLoop
		}
	}
}
