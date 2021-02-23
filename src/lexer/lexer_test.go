package lexer

import (
	"io/ioutil"
	"log"
	"testing"
	"unicode"
)

func TestLexer(t *testing.T) {
	t.Logf("\n%v\n", "moin\U00010348\uD55C\u20AC\a\x01\x03\x48")
	t.Logf("\n%v\n", rune(0x10348))
	t.Logf("\n%v\n", unicode.Is(unicode.Hex_Digit, '\\'))
	t.Logf("\n%v\n", []byte{'a', 'b', 'z', 'A', 'B', 'Z'})
	t.Logf("\n%v\n", string([]byte{64, 65, 66, 90, 91, 92, 93, 94, 95, 96, 97, 98, 122, 123}))

	buff, err := ioutil.ReadFile("../test/test.gosh")
	if err != nil {
		t.Fatal("Could not read ../test/test.gosh")
	}

	runArr := []rune(string(buff))
	lex := NewLexer(runArr, len(runArr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Logf("\n%v\n", *tokens)
}

func TestUnknownToken(t *testing.T) {
	buf, err := ioutil.ReadFile("../test/unknownToken.gosh")
	if err != nil {
		t.Fatal(err)
	}

	contents := []rune(string(buf))
	lex := NewLexer(contents, len(contents), "../test/unknownToken.gosh")
	_, lerr := lex.Lex()

	_, ok := lerr.(*UnknownTokenError)

	t.Log(lerr.Error())

	if !ok {
		t.Fatal("Expected an unknown error token, got", lerr.Error(), "instead")
	}
}

func TestMissingQuote(t *testing.T) {
	buf, err := ioutil.ReadFile("../test/missingQuote.gosh")
	if err != nil {
		t.Fatal(err)
	}

	contents := []rune(string(buf))

	lex := NewLexer(contents, len(contents), "../test/missingQuote.gosh")
	_, lerr := lex.Lex()

	_, ok := lerr.(*MissingQuoteError)

	t.Log(lerr.Error())

	if !ok {
		t.Fatal("Expected a missing quote error, got", lerr.Error(), "instead")
	}
}
