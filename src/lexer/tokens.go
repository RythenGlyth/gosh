package lexer

import (
	"fmt"
	"strings"
)

// TokenType is a type of token
type TokenType uint8

const (
	// if nothing was found
	TtEmpty TokenType = iota
	TtIdentifier
	// variables identifier need to be of the following form: [$§][A-Za-z_]+[A-z0-9_]*
	// public variables identifier, starting with $ (until space)
	TtPubVarIdent
	// private variables identifier, starting with § (until space)
	TtPrivVarIdent
	// string surrounded by quotes
	TtString
	// number in decimal, could also be in base 2 (0b) or in hex (0x)
	TtNumber

	TtTrue
	TtFalse
	TtIf
	TtElse
	TtFor
	TtReturn

	TtLParen
	TtRParen
	// {: start of block
	TtLBrace
	// }: end of block
	TtRBrace
	// [: start of array
	TtLBracket
	// ]: end of array
	TtRBracket
	TtSemicolon
	TtComma
	TtDot
	// Conditional expression
	TtQuestion
	// Advanced Forloop or conditional expression
	TtColon

	// Modulo
	TtPercent
	// multiplication
	TtStar
	// division
	TtSlash
	TtPlus
	TtMinus
	TtBang
	// increase
	TtPlusPlus
	// decrease
	TtMinusMinus
	// And
	TtAndAnd
	// ||: Or
	TtBarbar
	// Equals
	TtEqEq
	// !=: Not Equals
	TtBangEq
	// Smaller or equals
	TtLtEq
	// Greater or equals
	TtGtEq
	// Sets variables
	TtEq
	TtLt
	TtGt
	// |: Pipe left into right
	TtBar
	// Two Commands async at a time
	TtAnd
	TtPlusEq
	TtMinusEq
	TtStarEq
	TtSlashEq
	TtPerEq
	// =>: arrow (for lambda)
	TtEqGt
	// ->: arrow (for lambda)
	TtMinusGt
)

// Token for the Lexer
type Token struct {
	TokenType TokenType
	StartPos  int
	EndPos    int
	Value     interface{}
}

func (Tt TokenType) String() string {
	switch Tt {
	case TtIdentifier:
		return "TtIdentifier"
	case TtPubVarIdent:
		return "TtPubVarIdent"
	case TtPrivVarIdent:
		return "TtPrivVarIdent"
	case TtString:
		return "TtString"
	case TtNumber:
		return "TtNumber"
	case TtTrue:
		return "TtTrue"
	case TtFalse:
		return "TtFalse"
	case TtIf:
		return "TtIf"
	case TtElse:
		return "TtElse"
	case TtFor:
		return "TtFor"
	case TtReturn:
		return "TtReturn"
	case TtLParen:
		return "TtLParen"
	case TtRParen:
		return "TtRParen"
	case TtLBrace:
		return "TtLBrace"
	case TtRBrace:
		return "TtRBrace"
	case TtLBracket:
		return "TtLBracket"
	case TtRBracket:
		return "TtRBracket"
	case TtSemicolon:
		return "TtSemicolon"
	case TtComma:
		return "TtComma"
	case TtDot:
		return "TtDot"
	case TtQuestion:
		return "TtQuestion"
	case TtColon:
		return "TtColon"
	case TtPercent:
		return "TtPercent"
	case TtStar:
		return "TtStar"
	case TtSlash:
		return "TtSlash"
	case TtPlus:
		return "TtPlus"
	case TtMinus:
		return "TtMinus"
	case TtBang:
		return "TtBang"
	case TtPlusPlus:
		return "TtPlusPlus"
	case TtMinusMinus:
		return "TtMinusMinus"
	case TtAndAnd:
		return "TtAndAnd"
	case TtBarbar:
		return "TtBarbar"
	case TtEqEq:
		return "TtEqEq"
	case TtBangEq:
		return "TtBangEq"
	case TtLtEq:
		return "TtLtEq"
	case TtGtEq:
		return "TtGteq"
	case TtEq:
		return "TtEq"
	case TtLt:
		return "TtLt"
	case TtGt:
		return "TtGt"
	case TtBar:
		return "TtBar"
	case TtAnd:
		return "TtAnd"
	case TtPlusEq:
		return "TtPlusEq"
	case TtMinusEq:
		return "TtMinusEq"
	case TtStarEq:
		return "TtStarEq"
	case TtSlashEq:
		return "TtSlashEq"
	case TtPerEq:
		return "TtPerEq"
	case TtEqGt:
		return "TtEqGt"
	case TtMinusGt:
		return "TtMinusGt"
	default:
		return "TtEmpty"
	}
}

var MappedIt = map[string]TokenType{
	"%":  TtPercent,
	"*":  TtStar,
	"/":  TtSlash,
	"+":  TtPlus,
	"-":  TtMinus,
	"!":  TtBang,
	"++": TtPlusPlus,
	"--": TtMinusMinus,
	"&&": TtAndAnd,
	"||": TtBarbar,
	"==": TtEqEq,
	"!=": TtBangEq,
	"<=": TtLtEq,
	">=": TtGtEq,
	"=":  TtEq,
	"<":  TtLt,
	">":  TtGt,
	"|":  TtBar,
	"&":  TtAnd,
	"+=": TtPlusEq,
	"-=": TtMinusEq,
	"*=": TtStarEq,
	"/=": TtSlashEq,
	"%=": TtPerEq,
	"=>": TtEqGt,
	"->": TtMinusGt,
}

func (t *Token) String() string {
	var builder strings.Builder
	t.StringifyIntoBuilder(&builder)
	return builder.String()
}

func (t *Token) StringifyIntoBuilder(builder *strings.Builder) {
	builder.WriteRune('{')
	builder.WriteString(t.TokenType.String())
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprint(t.StartPos))
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprint(t.EndPos))
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprintf("%v", t.Value))
	builder.WriteRune('}')
}

func TokenArrayToString(tokenArray *[]Token) string {
	var builder strings.Builder
	builder.WriteRune('{')
	for _, t := range *tokenArray {
		t.StringifyIntoBuilder(&builder)
		builder.WriteRune(',')
		builder.WriteRune('\n')
	}
	builder.WriteRune('}')
	return builder.String()
}
