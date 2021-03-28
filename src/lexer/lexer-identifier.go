package lexer

import (
	"strings"
	"unicode"
)

func (lex *Lexer) readIdentifier() (identifier string, err LexError) {
	var identifierBuilder strings.Builder
	identifierBuilder.WriteRune(lex.character)
	for lex.position+1 < lex.length && !unicode.IsSpace(lex.buffer[lex.position+1]) && !unicode.Is(SpecialRangeTable, lex.buffer[lex.position+1]) {
		stringifiedIdentifierBuilder := identifierBuilder.String()

		if lex.buffer[lex.position+1] == '(' && (stringifiedIdentifierBuilder == "if" || stringifiedIdentifierBuilder == "for" || stringifiedIdentifierBuilder == "while") {
			return stringifiedIdentifierBuilder, nil
		}
		lex.next()
		identifierBuilder.WriteRune(lex.character)
	}
	return identifierBuilder.String(), nil
}

// SpecialRangeTable is
var SpecialRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{},
}
