package lexer

import (
	"strings"
	"unicode"
)

func (lex *Lexer) readVariableIdentifier(varIdentifierBuilder *strings.Builder) (err LexError) {
	//skip first character
	for lex.position+1 < lex.length && unicode.Is(VariableIdentifierRangeTable, lex.buffer[lex.position+1]) {
		lex.next()
		varIdentifierBuilder.WriteRune(lex.character)
	}
	return nil
}

// VariableIdentifierRangeTable is a set of all characters allowed in variable identifiers/names
var VariableIdentifierRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{'0', '9', 1},
		{'A', 'Z', 1},
		{'_', '_', 1},
		{'a', 'z', 1},
	},
}
