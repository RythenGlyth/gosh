package lexer

import (
	"strconv"
	"strings"
	"unicode"
)

// Lexer (Tokenizer) is created from an input (slice of runes) and separates it into tokens
// by callintg the Lex() function.
type Lexer struct {

	// buffer containing the character of the code to analyze
	buffer []rune

	// character inden in buffer
	position int

	codeXPos int
	codeYPos int

	// length of the content in the buffer
	length int

	inputName string

	character rune

	// if the last character was a backslash. used to handle escaped codes
	wasBackslash bool
}

// NewLexer creates a new Lexer
func NewLexer(source []rune, length int, inputName string) *Lexer {
	return &Lexer{source, -1, 0, 1, length, inputName, '\x00', false}
}

// Lex into tokens
func (lex *Lexer) Lex() (*[]Token, LexError) {
	var tokens []Token
	lex.next()

	for lex.position < lex.length {
		token, err := lex.nexTtoken()
		if err != nil {
			return &tokens, err
		}

		if token.TokenType != TtEmpty {
			tokens = append(tokens, *token)
		}
	}

	return &tokens, nil
}

func (lex *Lexer) nexTtoken() (*Token, LexError) {
	var startPos int = lex.position
	var tokenType TokenType = TtEmpty
	var valueBuilder strings.Builder

loop:
	for {
		if lex.position >= lex.length {
			break loop
		}
		if unicode.IsSpace(lex.character) {
			if lex.character == '\n' {
				lex.codeXPos = 0
				lex.codeYPos++
				tokenType = TtSemicolon
				break loop
			}
			lex.next()
			continue
		}
		switch lex.character {
		case '(':
			tokenType = TtLParen
			break loop
		case ')':
			tokenType = TtRParen
			break loop
		case ',':
			tokenType = TtComma
			break loop
		case '.':
			if lex.position+1 < lex.length {
				if unicode.IsNumber(lex.buffer[lex.position+1]) {
					var numVal float64
					lex.addDecimalNumber(10, &numVal)

					lex.next()

					var endpos int = lex.position
					return &Token{TtNumber, startPos, endpos, numVal}, nil
				}
			}
			tokenType = TtDot
			break loop
		case ':':
			tokenType = TtColon
			break loop
		case ';':
			tokenType = TtSemicolon
			break loop
		case '?':
			tokenType = TtQuestion
			break loop
		case '[':
			tokenType = TtLBracket
			break loop
		case ']':
			tokenType = TtRBracket
			break loop
		case '{':
			tokenType = TtLBrace
			break loop
		case '}':
			tokenType = TtRBrace
			break loop
		case '"', '\'':
			stringQuotes := lex.character
			lex.next()
			for lex.character != stringQuotes || lex.wasBackslash {
				if lex.position+1 < lex.length {
					valueBuilder.WriteRune(lex.character)
					lex.next()
				} else {
					return nil, &MissingQuoteError{Position{lex.codeXPos + 1, lex.codeYPos, lex.position + 1, lex}}
				}
			}
			tokenType = TtString
			break loop
		case '0':
			if lex.position+1 < lex.length {
				switch lex.buffer[lex.position+1] {
				case 'b', 'B':
					lex.next()
					lex.next()
					val, err := lex.readNumber(2)
					if err != nil {
						return nil, err
					}
					tokenType = TtNumber

					lex.next()

					var endpos int = lex.position
					return &Token{tokenType, startPos, endpos, val}, nil
				case 'x', 'X':
					lex.next()
					lex.next()
					val, err := lex.readNumber(16)
					if err != nil {
						return nil, err
					}
					tokenType = TtNumber

					lex.next()

					var endpos int = lex.position
					return &Token{tokenType, startPos, endpos, val}, nil
				// 0rn:hhh... (number with specific radix) n=radix (between 2 and 36), h=number itself
				case 'r', 'R':
					lex.next()
					var numStringBuilder strings.Builder
					for lex.position+1 < lex.length && lex.buffer[lex.position+1] != ':' {
						lex.next()
						numStringBuilder.WriteRune(lex.character)
					}
					lex.next()
					lex.next()
					radix, err := strconv.ParseFloat(numStringBuilder.String(), 32)
					if err != nil {
						return nil, &NumberFormatError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}}
					}
					val, err2 := lex.readNumber(radix)
					if err2 != nil {
						return nil, err2
					}
					tokenType = TtNumber

					lex.next()

					var endpos int = lex.position
					return &Token{tokenType, startPos, endpos, val}, nil
				}
			}
			fallthrough
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := lex.readNumber(10)
			if err != nil {
				return nil, err
			}
			tokenType = TtNumber

			lex.next()

			var endpos int = lex.position
			return &Token{tokenType, startPos, endpos, val}, nil
		case '$':
			err := lex.readVariableIdentifier(&valueBuilder)
			if err != nil {
				return nil, err
			}
			tokenType = TtPubVarIdent
			break loop
		case '§':
			err := lex.readVariableIdentifier(&valueBuilder)
			if err != nil {
				return nil, err
			}
			tokenType = TtPrivVarIdent
			break loop
		case '/':
			if lex.position+1 < lex.length {
				if lex.buffer[lex.position+1] == '*' {
					lex.nextPos()
					for lex.position+2 < lex.length && lex.buffer[lex.position+1] != '*' && lex.buffer[lex.position+2] != '/' {
						lex.nextPos()
					}
					lex.nextPos()
					lex.nextPos()
					break loop
				} else if lex.buffer[lex.position+1] == '/' {
					lex.nextPos()
					for lex.position+1 < lex.length && lex.buffer[lex.position+1] != '\n' {
						lex.nextPos()
					}
					lex.nextPos()
					break loop
				}
			}
			fallthrough
		case '%', '*', '+', '-', '|', '&', '=', '<', '>', '!':
			Tt, ok := MappedIt[string(lex.character)+string(lex.buffer[lex.position+1])]
			if lex.position+1 < lex.length && ok {
				tokenType = Tt
				lex.next()
				break loop
			} else {
				Tt, ok = MappedIt[string(lex.character)]
				if ok {
					tokenType = Tt
					break loop
				}
			}
			fallthrough
		default:
			identifier, err := lex.readIdentifier()
			if err != nil {
				return nil, err
			}
			switch strings.ToLower(identifier) {
			case "if":
				tokenType = TtIf
				valueBuilder.WriteString("if")
			case "else":
				tokenType = TtElse
				valueBuilder.WriteString("else")
			case "for", "while":
				tokenType = TtFor
				valueBuilder.WriteString("for")
			case "return":
				tokenType = TtReturn
				valueBuilder.WriteString("return")
			case "true":
				tokenType = TtTrue
				valueBuilder.WriteString("true")
			case "false":
				tokenType = TtFalse
				valueBuilder.WriteString("false")
			case "do", "loop", "is", "try", "catch", "run", "switch", "case", "break", "continue", "register", "goto":
				return nil, &ReservedIdentifierError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}, identifier}
			default:
				tokenType = TtString
				valueBuilder.WriteString(identifier)
			}
			break loop
			//return nil, &UnknownTokenError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}}
		}
	}

	lex.next()

	var endpos int = lex.position
	return &Token{tokenType, startPos, endpos, valueBuilder.String()}, nil
}

func (lex *Lexer) next() LexError {
	return lex.nextB(false)
}

func (lex *Lexer) nextPos() {
	lex.position++
	lex.codeXPos++
}

// BinaryRangeTable is a set of 0 and 1
var BinaryRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0030, 0x0031, 1},
	},
}

// DecimalRangeTable is a set of all decimals
var DecimalRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '9', 1},
	},
}
