package parser

import (
	"gosh/src/lexer"
	"io/ioutil"
	"os"
	"testing"
)

func TestBlockStatement(t *testing.T) {
	r := []rune(`
for() {
	if() {
		;;;;;
		for({};;) {
			$x = 
		}
	}
	;;
}
`)

	lex := lexer.NewLexer(r, len(r), "f")

	tokens, lerr := lex.Lex()

	if lerr != nil {
		t.Fatal(lerr.Error())
	}
	t.Log(lexer.TokenArrayToString(tokens))

	p := NewParser(*tokens)

	s, err := p.Parse()

	if err != nil {
		t.Fatal(err)
	}

	s.Debug(os.Stdout, 0, "")

}

func TestParser(t *testing.T) {
	buf, err := ioutil.ReadFile("../test/test.gosh")
	if err != nil {
		t.Fatal("Could not read ../test/test.gosh")
	}

	r := []rune(string(buf))
	lex := lexer.NewLexer(r, len(r), "../test/test.gosh")

	tokens, lerr := lex.Lex()

	if lerr != nil {
		t.Fatal(lerr.Error())
	}

	p := NewParser(*tokens)

	p.Parse()
}
