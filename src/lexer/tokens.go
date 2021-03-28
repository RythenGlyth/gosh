package lexer

import (
	"fmt"
	"strings"
)

// TokenType is a type of token
type TokenType uint8

const (
	// if nothing was found
	ttEmpty TokenType = iota
	ttIdentifier
	// variables identifier need to be of the following form: [$§][A-Za-z_]+[A-z0-9_]*
	// public variables identifier, starting with $ (until space)
	ttPubVarIdent
	// private variables identifier, starting with § (until space)
	ttPrivVarIdent
	// string surrounded by quotes
	ttString
	// number in decimal, could also be in base 2 (0b) or in hex (0x)
	ttNumber

	ttTrue
	ttFalse
	ttIf
	ttElse
	ttFor
	ttReturn

	ttLParen
	ttRParen
	// {: start of block
	ttLBrace
	// }: end of block
	ttRBrace
	// [: start of array
	ttLBracket
	// ]: end of array
	ttRBracket
	ttSemicolon
	ttComma
	ttDot
	// Conditional expression
	ttQuestion
	// Advanced Forloop or conditional expression
	ttColon

	// Modulo
	ttPercent
	// multiplication
	ttStar
	// division
	ttSlash
	// Line Comment
	ttSlashSlash
	// Start of Block Comment
	ttSlashStar
	// End of Block Comment
	ttStarSlash
	ttPlus
	ttMinus
	ttBang
	// increase
	ttPlusPlus
	// decrease
	ttMinusMinus
	// And
	ttAndAnd
	// ||: Or
	ttBarbar
	// Equals
	ttEqEq
	// !=: Not Equals
	ttBangEq
	// Smaller or equals
	ttLtEq
	// Greater or equals
	ttGtEq
	// Sets variables
	ttEq
	ttLt
	ttGt
	// |: Pipe left into right
	ttBar
	// Two Commands async at a time
	ttAnd
	ttPlusEq
	ttMinusEq
	ttStarEq
	ttSlashEq
	ttPerEq
	// =>: arrow (for lambda)
	ttEqGt
	// ->: arrow (for lambda)
	ttMinusGt
)

// Token for the Lexer
type Token struct {
	tokenType TokenType
	startPos  int
	endPos    int
	value     interface{}
}

func (tt TokenType) String() string {
	switch tt {
	case ttIdentifier:
		return "ttIdentifier"
	case ttPubVarIdent:
		return "ttPubVarIdent"
	case ttPrivVarIdent:
		return "ttPrivVarIdent"
	case ttString:
		return "ttString"
	case ttNumber:
		return "ttNumber"
	case ttTrue:
		return "ttTrue"
	case ttFalse:
		return "ttFalse"
	case ttIf:
		return "ttIf"
	case ttElse:
		return "ttElse"
	case ttFor:
		return "ttFor"
	case ttReturn:
		return "ttReturn"
	case ttLParen:
		return "ttLParen"
	case ttRParen:
		return "ttRParen"
	case ttLBrace:
		return "ttLBrace"
	case ttRBrace:
		return "ttRBrace"
	case ttLBracket:
		return "ttLBracket"
	case ttRBracket:
		return "ttRBracket"
	case ttSemicolon:
		return "ttSemicolon"
	case ttComma:
		return "ttComma"
	case ttDot:
		return "ttDot"
	case ttQuestion:
		return "ttQuestion"
	case ttColon:
		return "ttColon"
	case ttPercent:
		return "ttPercent"
	case ttStar:
		return "ttStar"
	case ttSlash:
		return "ttSlash"
	case ttSlashSlash:
		return "ttSlashSlash"
	case ttSlashStar:
		return "ttSlashStar"
	case ttStarSlash:
		return "ttStarSlash"
	case ttPlus:
		return "ttPlus"
	case ttMinus:
		return "ttMinus"
	case ttBang:
		return "ttBang"
	case ttPlusPlus:
		return "ttPlusPlus"
	case ttMinusMinus:
		return "ttMinusMinus"
	case ttAndAnd:
		return "ttAndAnd"
	case ttBarbar:
		return "ttBarbar"
	case ttEqEq:
		return "ttEqEq"
	case ttBangEq:
		return "ttBangEq"
	case ttLtEq:
		return "ttLtEq"
	case ttGtEq:
		return "ttGteq"
	case ttEq:
		return "ttEq"
	case ttLt:
		return "ttLt"
	case ttGt:
		return "ttGt"
	case ttBar:
		return "ttBar"
	case ttAnd:
		return "ttAnd"
	case ttPlusEq:
		return "ttPlusEq"
	case ttMinusEq:
		return "ttMinusEq"
	case ttStarEq:
		return "ttStarEq"
	case ttSlashEq:
		return "ttSlashEq"
	case ttPerEq:
		return "ttPerEq"
	case ttEqGt:
		return "ttEqGt"
	case ttMinusGt:
		return "ttMinusGt"
	default:
		return "ttEmpty"
	}
}

var MappedIt = map[string]TokenType{
	"%":  ttPercent,
	"*":  ttStar,
	"/":  ttSlash,
	"//": ttSlashSlash,
	"/*": ttSlashStar,
	"*/": ttStarSlash,
	"+":  ttPlus,
	"-":  ttMinus,
	"!":  ttBang,
	"++": ttPlusPlus,
	"--": ttMinusMinus,
	"&&": ttAndAnd,
	"||": ttBarbar,
	"==": ttEqEq,
	"!=": ttBangEq,
	"<=": ttLtEq,
	">=": ttGtEq,
	"=":  ttEq,
	"<":  ttLt,
	">":  ttGt,
	"|":  ttBar,
	"&":  ttAnd,
	"+=": ttPlusEq,
	"-=": ttMinusEq,
	"*=": ttStarEq,
	"/=": ttSlashEq,
	"%=": ttPerEq,
	"=>": ttEqGt,
	"->": ttMinusGt,
}

func (t *Token) String() string {
	var builder strings.Builder
	t.StringifyIntoBuilder(&builder)
	return builder.String()
}

func (t *Token) StringifyIntoBuilder(builder *strings.Builder) {
	builder.WriteRune('{')
	builder.WriteString(t.tokenType.String())
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprint(t.startPos))
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprint(t.endPos))
	builder.WriteRune(',')
	builder.WriteRune(' ')
	builder.WriteString(fmt.Sprintf("%v", t.value))
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
