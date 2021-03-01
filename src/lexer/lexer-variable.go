package lexer

import (
	"strings"
	"unicode"
)

func (lex *Lexer) readVariableIdentifier() (varName string, err LexError) {
	var varIdentifierBuilder strings.Builder
	for unicode.Is(SingleSpecialRangeTable, lex.buffer[lex.position+1]) {
		lex.next()
		varIdentifierBuilder.WriteRune(lex.character)
	}
	return varIdentifierBuilder.String(), nil
}

// VariableIdentifierRangeTable is a set of all characters allowed in variable identifiers/names
var VariableIdentifierRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'', ')', 1},
	},
}
