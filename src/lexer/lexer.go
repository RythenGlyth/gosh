package lexer

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Lexer (Tokenizer)
type Lexer struct {
	//buffer containing the character of the code to analyze
	buffer []rune
	//character inden in buffer
	position int

	codeXPos int
	codeYPos int
	//length of the content in the buffer
	length int
	//if the last character was a backslash. used to handle escaped codes
	wasBackslash bool

	inputName string

	character rune
}

// NewLexer creates a new Lexer
func NewLexer(source []rune, length int, inputName string) *Lexer {
	return &Lexer{source, -1, 1, 1, length, false, inputName, '\x00'}
}

// Lex into tokens
func (lex *Lexer) Lex() (*[]Token, *LexError) {
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

func (lex *Lexer) nextToken() (*Token, *LexError) {
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
		case ';', ',', '.', '(', ')', '[', ']', '{', '}':
			tokenType = TokenType(string(lex.character))
			break loop
		case '"', '\'':
			stringQuotes := lex.character
			lex.next()
			for lex.character != stringQuotes || lex.wasBackslash {
				if lex.position+1 < lex.length {
					valueBuilder.WriteRune(lex.character)
					lex.next()
				} else {
					return nil, &LexError{errors.New("missing closing quotes of string"), lex.codeXPos, lex.codeYPos, lex.position, lex}
				}
			}
			tokenType = ttString
			fmt.Print("sos")
			break loop
		case '0':
			if lex.position+1 < lex.length {
				switch lex.buffer[lex.position+1] {
				case 'b', 'B':
					lex.next()
					/*var binaryStringBuilder strings.Builder
					for lex.position+1 < lex.length && unicode.Is(BinaryRangeTable, lex.buffer[lex.position+1]) {
						lex.next()
						binaryStringBuilder.WriteRune(lex.character)
					}
					parsed, err := strconv.ParseInt(binaryStringBuilder.String(), 2, 64)
					if err != nil {
						return nil, &LexError{errors.New("Cannot read binary: " + err.Error()), lex.codeXPos, lex.codeYPos, lex.position, lex}
					}
					tokenType = ttNumber
					valueBuilder.WriteString(fmt.Sprint(parsed))
					break loop*/
					val, err := lex.readNumber(2)
					if err != nil {
						return nil, err
					}
					tokenType = ttNumber
					valueBuilder.WriteString(fmt.Sprint(val))
					break loop
				case 'x', 'X':
					lex.next()
					/*var hexStringBuilder strings.Builder
					for lex.position+1 < lex.length && unicode.Is(unicode.Hex_Digit, lex.buffer[lex.position+1]) {
						lex.next()
						hexStringBuilder.WriteRune(lex.character)
					}
					parsed, err := strconv.ParseInt(hexStringBuilder.String(), 16, 64)
					if err != nil {
						return nil, &LexError{errors.New("Cannot read hex: " + err.Error()), lex.codeXPos, lex.codeYPos, lex.position, lex}
					}
					tokenType = ttNumber
					valueBuilder.WriteString(fmt.Sprint(parsed))
					break loop*/
					val, err := lex.readNumber(16)
					if err != nil {
						return nil, err
					}
					tokenType = ttNumber
					valueBuilder.WriteString(fmt.Sprint(val))
					break loop
				//0rn:hhh... (number with specific radix) n=radix (between 2 and 36), h=number itself
				case 'r', 'R':
					lex.next()
					var numStringBuilder strings.Builder
					for lex.position+1 < lex.length && lex.buffer[lex.position+1] != ':' {
						lex.next()
						numStringBuilder.WriteRune(lex.character)
					}
					lex.next()
					radix, err := strconv.ParseFloat(numStringBuilder.String(), 10)
					if err != nil {
						return nil, &LexError{errors.New("Wrong format (it should be 0rnn:hh...) "), lex.codeXPos, lex.codeYPos, lex.position, lex}
					}
					val, err2 := lex.readNumber(radix)
					if err2 != nil {
						return nil, err2
					}
					tokenType = ttNumber
					valueBuilder.WriteString(fmt.Sprint(val))
					break loop
				}
			}
			fallthrough
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := lex.readNumber(10)
			if err != nil {
				return nil, err
			}
			tokenType = ttNumber
			valueBuilder.WriteString(fmt.Sprint(val))
			break loop
		default:
			return nil, &LexError{errors.New("Unknown Token"), lex.codeXPos, lex.codeYPos, lex.position, lex}
			//fmt.Print("sas")
		}
	}
	lex.next()
	var endpos int = lex.position
	return &Token{tokenType, startPos, endpos, valueBuilder.String()}, nil
}

func (lex *Lexer) next() *LexError {
	return lex.nextB(false)
}

func (lex *Lexer) nextB(wasBackslash bool) *LexError {
	lex.wasBackslash = wasBackslash
	lex.nextPos()
	if lex.position >= 0 && lex.position < lex.length && lex.buffer[lex.position] == '\\' && !lex.wasBackslash {
		lex.wasBackslash = true
		if lex.position+1 < lex.length {
			return lex.nextB(true)
		} else {
			return &LexError{errors.New("nothing to escape"), lex.codeXPos, lex.codeYPos, lex.position, lex}
		}
	}
	if lex.position >= 0 && lex.position < lex.length {
		if lex.wasBackslash {
			switch lex.buffer[lex.position] {
			case 'b':
				lex.character = '\b'
			case 'f':
				lex.character = '\f'
			case 'n':
				lex.character = '\n'
			case 'r':
				lex.character = '\r'
			case 't':
				lex.character = '\t'
			case 'v':
				lex.character = '\v'
			case '\\':
				lex.character = '\\'
			case '\'':
				lex.character = '\''
			case '"':
				lex.character = '"'
			case 'x', 'u', 'U':
				var hexStringBuilder strings.Builder
				hexStringBuilder.WriteRune(lex.buffer[lex.position+1])
				lex.nextPos()
				for unicode.Is(unicode.Hex_Digit, lex.buffer[lex.position+1]) {
					hexStringBuilder.WriteRune(lex.buffer[lex.position+1])
					lex.nextPos()
				}
				var err error
				var decoded int64
				hexString := hexStringBuilder.String()
				decoded, err = strconv.ParseInt(hexString, 16, 64)
				if err != nil {
					return &LexError{errors.New("can't decode " + hexString + " in hex format"), lex.codeXPos, lex.codeYPos, lex.position, lex}
				}
				lex.character = rune(decoded)
			default:
				return &LexError{errors.New("can't escape " + string(lex.buffer[lex.position])), lex.codeXPos, lex.codeYPos, lex.position, lex}
			}
		} else {
			lex.character = lex.buffer[lex.position]
		}
	}
	return nil
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

func (lex *Lexer) readNumber(radix float64) (float64, *LexError) {
	var currentVal float64
loop:
	for {
		if lex.position+1 >= lex.length {
			break loop
		}
		if lex.buffer[lex.position+1] == '.' {
			lex.next()
			var decVal float64
		decLoop:
			for {
				var thisVal = getNumberValue(lex.buffer[lex.position+1])
				if thisVal > 0 && thisVal < radix {

				} else {
					break decLoop
				}
			}
			currentVal += decVal
			break loop
		} else {
			currentVal *= radix
			var thisVal = getNumberValue(lex.buffer[lex.position+1])
			fmt.Print("_" + string(lex.buffer[lex.position+1]) + "_" + fmt.Sprint(thisVal) + ": " + fmt.Sprint(thisVal) + "\n")
			if thisVal >= 0 && thisVal < radix {
				currentVal += thisVal
				lex.next()
			} else {
				break loop
			}
		}
	}
	return currentVal, nil
}

func getNumberValue(char rune) float64 {
	if char >= '0' && char <= '9' {
		return (float64)(char - '0')
	}
	if char >= 'A' && char <= 'Z' {
		return (float64)(char - 'A' + 10)
	}
	if char >= 'a' && char <= 'z' {
		return (float64)(char - 'a' + 10)
	}
	return -1
}

//LexError is the type of error thrown by the lexer
type LexError struct {
	err      error
	codeXPos int
	codeYPos int
	bufPos   int
	lexer    *Lexer
}

//SPrint s the lexer error into a string
func (lexErr *LexError) SPrint() string {
	return ("Error while Lexing:" + "\n" + lexErr.lexer.inputName + ":" + fmt.Sprint(lexErr.codeYPos) + ":" + fmt.Sprint(lexErr.codeXPos) + ": " + lexErr.err.Error() + "\n" + string(lexErr.lexer.buffer[int(math.Max(float64(lexErr.bufPos-10), float64(0))):int(math.Min(float64(lexErr.bufPos+10), float64(lexErr.lexer.length-1)))]) + "\n" + strings.Repeat(" ", (lexErr.bufPos-int(math.Max(float64(lexErr.bufPos-10), float64(0))))) + "^")
}

//BinaryRangeTable is a set of 0 and 1
var BinaryRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0030, 0x0031, 1},
	},
}

//DecimalRangeTable is a set of all decimals
var DecimalRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '9', 1},
	},
}
