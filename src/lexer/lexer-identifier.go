package lexer

import (
	"strings"
	"unicode"
)

func (lex *Lexer) readIdentifier() (identifier string, err LexError) {
	var identifierBuilder strings.Builder
	for !unicode.IsSpace(lex.buffer[lex.position+1]) && !unicode.Is(SpecialRangeTable, lex.buffer[lex.position+1]) {
		lex.next()
		identifierBuilder.WriteRune(lex.character)
	}
	return identifierBuilder.String(), nil
}
