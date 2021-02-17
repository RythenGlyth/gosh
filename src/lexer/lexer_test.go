package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	t.Logf("\n%v\n", []byte{'a', 'b', 'z', 'A', 'B', 'Z'})
	t.Logf("\n%v\n", string([]byte{64, 65, 66, 90, 91, 92, 93, 94, 95, 96, 97, 98, 122, 123}))
	var lex *Lexer
	var buff []byte = []byte("...;;")
	lex = NewLexer(buff, len(buff))
	tokens := lex.Lex()
	t.Logf("\n%v\n", tokens)
}
