package lexer

// TokenType is
type TokenType string

const (
	//if nothing was found
	ttEmpty TokenType = "empty"
	//string without quotes, could be a string or things like if or for, need to be of the following form: [A-z_][A-z0-9_\-/]* (can only start with  (until space)
	ttIdentifier TokenType = "identifier"
	//variables identifier need to be of the following form: [$§][A-z_]+[A-z0-9_]*
	//public variables identifier, starting with $ (until space)
	ttPubVarIdent TokenType = "pubVarIdent"
	//private variables identifier, starting with § (until space)
	ttPrivVarIdent TokenType = "privVarIdent"
	//string surrounded by quotes
	ttString TokenType = "string"
	//number in decimal, could also be in base 2 (0b) or in hex (0x)
	ttNumber TokenType = "number"

	ttLParen    TokenType = "("
	ttRParen    TokenType = ")"
	ttLBrace    TokenType = "{"
	ttRBrace    TokenType = "}"
	ttLBracket  TokenType = "["
	ttRBracket  TokenType = "]"
	ttSemicolon TokenType = ";"
	ttComma     TokenType = ","
	ttDot       TokenType = "."
	//Modulo
	ttPercent TokenType = "%"
	//multiplication
	ttStar TokenType = "*"

	//division
	ttSlash TokenType = "/"
	//Line Comment
	ttSlashSlash TokenType = "//"
	//Start of Block Comment
	ttSlashStar TokenType = "/*"
	//End of Block Comment
	ttStarSlash TokenType = "*/"
	ttPlus      TokenType = "+"
	ttMinus     TokenType = "-"
	ttBang      TokenType = "!"
	//increase
	ttPlusPlus TokenType = "++"
	//decrease
	ttMinusMinus TokenType = "--"
	//And
	ttAndAnd TokenType = "&&"
	//Or
	ttBarbar TokenType = "||"
	//Equals
	ttEqEq TokenType = "=="
	//Not Equals
	ttBangEq TokenType = "!="
	//Smaller or equals
	ttLtEq TokenType = "<="
	//Greater or equals
	ttGteq TokenType = ">="
	//Sets variables
	ttEq TokenType = "="
	ttLt TokenType = "<"
	ttGt TokenType = ">"
	//Pipe left into right
	ttBar TokenType = "|"
	//Two Commands async at a time
	ttAnd TokenType = "&"
)

//Token for the Lexer
type Token struct {
	tokenType TokenType
	startPos  int
	endPos    int
}
