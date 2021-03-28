package lexer

import (
	"io/ioutil"
	"log"
	"reflect"
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

	t.Log(TokenArrayToString(tokens))
}

func TestMissingQuote(t *testing.T) {
	buf, err := ioutil.ReadFile("../test/missingQuote.gosh")
	if err != nil {
		t.Fatal(err)
	}

	contents := []rune(string(buf))

	lex := NewLexer(contents, len(contents), "../test/missingQuote.gosh")
	_, lerr := lex.Lex()

	if lerr != nil {
		_, ok := lerr.(*MissingQuoteError)

		if !ok {
			t.Fatal("Expected a missing quote error, got", lerr.Error(), "instead")
		}
	} else {
		t.Fatal("Expected a missing quote error, got no error")
	}
}

func TestVar(t *testing.T) {
	arr := []rune("§test23er $789test_01")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Logf("\n%v\n", *tokens)

	if !tokenArrayEqual(*tokens, []Token{{ttPrivVarIdent, 0, 9, "test23er"}, {ttPubVarIdent, 9, 21, "789test_01"}}) {
		t.FailNow()
	}
}

func tokenArrayEqual(a, b []Token) bool {
	return reflect.DeepEqual(a, b)
}

func TestIdentifier(t *testing.T) {
	arr := []rune("if() { sos nä? ife }")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Logf("\n%v\n", *tokens)

	if !tokenArrayEqual(*tokens, []Token{{ttIf, 0, 2, "if"}, {ttLParen, 2, 3, ""}, {ttRParen, 3, 4, ""}, {ttLBrace, 4, 6, ""}, {ttString, 6, 10, "sos"}, {ttString, 10, 14, "nä?"}, {ttString, 14, 18, "ife"}, {ttRBrace, 18, 20, ""}}) {
		t.FailNow()
	}
}

func TestNumerus(t *testing.T) {

	arr := []rune("3 0x410 0b101 3.1 7.54 3.01 0xf.8 3.000000000000001")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Log(TokenArrayToString(tokens))

	if !tokenArrayEqual(*tokens, []Token{{ttIf, 0, 2, "if"}, {ttLParen, 2, 3, ""}, {ttRParen, 3, 4, ""}, {ttLBrace, 4, 6, ""}, {ttString, 6, 10, "sos"}, {ttString, 10, 14, "nä?"}, {ttString, 14, 18, "ife"}, {ttRBrace, 18, 20, ""}}) {
		t.FailNow()
	}
}
