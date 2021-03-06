package lexer

import (
	"encoding/hex"
	"fmt"
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
		token, err := lex.nextToken()
		if err != nil {
			return &tokens, err
		}

		if token.tokenType != ttEmpty {
			tokens = append(tokens, *token)
		}
	}

	return &tokens, nil
}

func (lex *Lexer) nextToken() (*Token, LexError) {
	var startPos int = lex.position
	var tokenType TokenType = ttEmpty
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
			}
			lex.next()
			continue
		}
		switch lex.character {
		case '(':
			tokenType = ttLParen
			break loop
		case ')':
			tokenType = ttRParen
			break loop
		case ',':
			tokenType = ttComma
			break loop
		case '.':
			if lex.position+1 < lex.length {
				if unicode.IsNumber(lex.buffer[lex.position+1]) {
					var numVal float64
					lex.addDecimalNumber(10, &numVal)

					lex.next()

					var endpos int = lex.position
					return &Token{ttNumber, startPos, endpos, numVal}, nil
				}
			}
			tokenType = ttDot
			break loop
		case ':':
			tokenType = ttColon
			break loop
		case ';':
			tokenType = ttSemicolon
			break loop
		case '?':
			tokenType = ttQuestion
			break loop
		case '[':
			tokenType = ttLBracket
			break loop
		case ']':
			tokenType = ttRBracket
			break loop
		case '{':
			tokenType = ttLBrace
			break loop
		case '}':
			tokenType = ttRBrace
			break loop
		case '"', '\'':
			stringQuotes := lex.character
			lex.next()
			for lex.character != stringQuotes || lex.wasBackslash {
				fmt.Print(string(lex.character))
				if lex.position+1 < lex.length {
					valueBuilder.WriteRune(lex.character)
					lex.next()
				} else {
					return nil, &MissingQuoteError{Position{lex.codeXPos + 1, lex.codeYPos, lex.position + 1, lex}}
				}
			}
			tokenType = ttString
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
					tokenType = ttNumber

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
					tokenType = ttNumber

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
					radix, err := strconv.ParseFloat(numStringBuilder.String(), 10)
					if err != nil {
						return nil, &NumberFormatError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}}
					}
					val, err2 := lex.readNumber(radix)
					if err2 != nil {
						return nil, err2
					}
					tokenType = ttNumber

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
			tokenType = ttNumber

			lex.next()

			var endpos int = lex.position
			return &Token{tokenType, startPos, endpos, val}, nil
		case '$':
			err := lex.readVariableIdentifier(&valueBuilder)
			if err != nil {
				return nil, err
			}
			tokenType = ttPubVarIdent
			break loop
		case '§':
			err := lex.readVariableIdentifier(&valueBuilder)
			if err != nil {
				return nil, err
			}
			tokenType = ttPrivVarIdent
			break loop
		case '%', '*', '/', '+', '-', '|', '&', '=', '<', '>', '!':
			tt, ok := MappedIt[string(lex.character)+string(lex.buffer[lex.position+1])]
			if lex.position+1 < lex.length && ok {
				tokenType = tt
				lex.next()
				break loop
			} else {
				tt, ok = MappedIt[string(lex.character)]
				if ok {
					tokenType = tt
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
				tokenType = ttIf
				valueBuilder.WriteString("if")
			case "else":
				tokenType = ttElse
				valueBuilder.WriteString("else")
			case "for", "while":
				tokenType = ttFor
				valueBuilder.WriteString("for")
			case "return":
				tokenType = ttReturn
				valueBuilder.WriteString("return")
			case "true":
				tokenType = ttTrue
				valueBuilder.WriteString("true")
			case "false":
				tokenType = ttFalse
				valueBuilder.WriteString("false")
			case "do", "loop", "is", "try", "catch", "run", "switch", "case", "break", "continue", "register", "goto":
				return nil, &ReservedIdentifierError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}, identifier}
			default:
				tokenType = ttString
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

	if lex.position >= 0 && lex.position < lex.length {
		fmt.Print(hex.EncodeToString([]byte(string(lex.buffer[lex.position : lex.position+1]))))
		fmt.Print(": " + strconv.FormatBool(lex.wasBackslash) + " : " + fmt.Sprint(lex.codeXPos))
		defer fmt.Print("\n")
	}
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
