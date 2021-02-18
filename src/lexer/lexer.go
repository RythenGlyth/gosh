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
	return &Lexer{source, -1, 1, 0, length, false, inputName, '\x00'}
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
		tokens = append(tokens, *token)
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
		switch lex.character {
		case ' ', '\t', '\f', '\r':
			lex.next()
		case '\n':
			lex.codeXPos = 0
			lex.codeYPos++
			lex.next()
		case ';', ',', '.', '(', ')', '[', ']', '{', '}', '%', '*':
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
		default:
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
			fmt.Print(" sas")
			fmt.Print(string(lex.buffer[lex.position]) + " ")
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
				hexString := string(lex.buffer[lex.position+1])
				lex.position++
				for unicode.Is(unicode.Hex_Digit, lex.buffer[lex.position+1]) {
					hexString += string(lex.buffer[lex.position+1])
					lex.position++
				}
				var err error
				var decoded int64
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
		fmt.Print(": " + strconv.FormatBool(lex.wasBackslash))
		defer fmt.Print("\n")
	}
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
	return ("Error while Lexing:" + "\n" + lexErr.lexer.inputName + ":" + fmt.Sprint(lexErr.codeYPos) + ":" + fmt.Sprint(lexErr.codeXPos) + ": " + lexErr.err.Error() + "\n" + string(lexErr.lexer.buffer[int(math.Max(float64(lexErr.bufPos-10), float64(0))):int(math.Min(float64(lexErr.bufPos+10), float64(lexErr.lexer.length-1)))]))
}
