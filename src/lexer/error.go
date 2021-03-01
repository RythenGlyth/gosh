package lexer

import (
	"fmt"
	"math"
	"strings"
)

// LexError is a generic error that occurs during lexical analysis.
type LexError interface {
	Error() string
}

// Position is the type of error thrown by the lexer.
type Position struct {
	codeXPos int
	codeYPos int
	bufPos   int
	lexer    *Lexer
}

func (p *Position) String() string {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("%s: l %d col %d\n", p.lexer.inputName, p.codeYPos, p.codeXPos))

	start := int(math.Max(float64(p.bufPos-10), float64(0))) // FIXME this ignores newlines
	end := int(math.Min(float64(p.bufPos+10), float64(p.lexer.length)))

	out.WriteString(string(p.lexer.buffer[start:end]))
	out.WriteString("\n")

	for i := 0; i < p.bufPos-int(math.Max(float64(p.bufPos-10), float64(0))); i++ {
		out.WriteString(" ")
	}

	out.WriteString("^")

	return out.String()
}

// UnknownTokenError is returned if an unknown token is encountered.
type UnknownTokenError struct {
	pos Position
}

func (e *UnknownTokenError) Error() string {
	return "unknown token at " + e.pos.String()
}

// MissingQuoteError is returned if a string hasn't been closed.
type MissingQuoteError struct {
	pos Position
}

func (e *MissingQuoteError) Error() string {
	return "missing closing quotes at " + e.pos.String()
}

// NumberFormatError is returned if the number is formatted wrong.
type NumberFormatError struct {
	pos Position
}

func (e *NumberFormatError) Error() string {
	return "wrong number format, should be 0rnn:hh... at " + e.pos.String()
}

// TrailingBackslashError is returned if there is a character missing after a backslash.
type TrailingBackslashError struct {
	pos Position
}

func (e *TrailingBackslashError) Error() string {
	return "trailing backslash error, nothing to escape at " + e.pos.String()
}

// InvalidEscapeCharacterError is returned if an invalid character is escaped.
type InvalidEscapeCharacterError struct {
	pos Position
	ch  rune
}

func (e *InvalidEscapeCharacterError) Error() string {
	return "can't escape " + string(e.ch) + " at " + e.pos.String()
}

// ReservedIdentifierError is returned if an identifier uses reserved keywords.
type ReservedIdentifierError struct {
	pos        Position
	identifier string
}

func (e *ReservedIdentifierError) Error() string {
	return "reserved identifier \"" + string(e.identifier) + "\" at " + e.pos.String()
}
