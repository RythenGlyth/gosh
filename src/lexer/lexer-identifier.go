package lexer

import (
	"strings"
	"unicode"
)

func (lex *Lexer) readIdentifier() (identifier string, err LexError) {
	var identifierBuilder strings.Builder
	for lex.position+1 < lex.length && !unicode.IsSpace(lex.buffer[lex.position+1]) && !unicode.Is(SpecialRangeTable, lex.buffer[lex.position+1]) {
		lex.next()
		identifierBuilder.WriteRune(lex.character)
	}
	return identifierBuilder.String(), nil
}

// SpecialRangeTable is
var SpecialRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'a', 'z', 1},
		{'A', 'Z', 1},
		{'0', '9', 1},
		{'_', '_', 1},
	},
	R32: []unicode.Range32{
		{'a', 'z', 1},
		{'A', 'Z', 1},
		{'0', '9', 1},
		{'_', '_', 1},
	},
}
