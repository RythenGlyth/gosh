package lexer

import (
	"fmt"
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

	if !tokenArrayEqual(*tokens, []Token{{TtPrivVarIdent, 0, 9, "test23er"}, {TtPubVarIdent, 9, 21, "789test_01"}}) {
		t.FailNow()
	}
}

func tokenArrayEqual(a, b []Token) bool {
	if len(a) != len(b) {
		fmt.Println("Lengths differ")
		return false
	}
	for i, ta := range a {
		tb := b[i]

		if ta.EndPos != tb.EndPos {
			fmt.Printf("endPos of %d (%v : %v) of differ\n", i, ta.String(), tb.String())
			return false
		}
		if ta.StartPos != tb.StartPos {
			fmt.Printf("startPos of %d (%v : %v) of differ\n", i, ta.String(), tb.String())
			return false
		}
		if ta.TokenType != tb.TokenType {
			fmt.Printf("tokenType of %d (%v : %v) of differ\n", i, ta.String(), tb.String())
			return false
		}
		if ta.Value != tb.Value {
			fmt.Printf("value of %d (%v : %v) of differ\n", i, ta.String(), tb.String())
			return false
		}
	}
	return true
}

func TestIdentifier(t *testing.T) {
	arr := []rune("if() { sos nä? ife }")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Logf("\n%v\n", *tokens)

	if !tokenArrayEqual(*tokens, []Token{{TtIf, 0, 2, "if"}, {TtLParen, 2, 3, ""}, {TtRParen, 3, 4, ""}, {TtLBrace, 4, 6, ""}, {TtString, 6, 10, "sos"}, {TtString, 10, 14, "nä?"}, {TtString, 14, 18, "ife"}, {TtRBrace, 18, 20, ""}}) {
		t.FailNow()
	}
}

func TestNumerus(t *testing.T) { //TODO .5

	arr := []rune("3 0x410 0b101 3.1 7.54 3.01 0xf.8 3.000000000000001 0r3:10 0xD55C .5 0x.1")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Log(TokenArrayToString(tokens))

	if !tokenArrayEqual(*tokens, []Token{{TtNumber, 0, 1, 3.0}, {TtNumber, 1, 7, 1040.0}, {TtNumber, 7, 13, 5.0}, {TtNumber, 13, 17, 3.1}, {TtNumber, 17, 22, 7.54}, {TtNumber, 22, 27, 3.01}, {TtNumber, 27, 33, 15.5}, {TtNumber, 33, 51, 3.000000000000001}, {TtNumber, 51, 58, 3.0}, {TtNumber, 58, 65, 54620.0}, {TtNumber, 65, 68, 0.5}, {TtNumber, 68, 73, 1.0 / 16}}) {
		t.FailNow()
	}
}

func TestSpecials(t *testing.T) {
	arr := []rune("\"\\uD55C\\U12323\"")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Log(TokenArrayToString(tokens))

	if !tokenArrayEqual(*tokens, []Token{{TtString, 0, 15, "\uD55C\U00012323"}}) {
		t.FailNow()
	}
}

func TestComments(t *testing.T) {
	arr := []rune("//This is a comment\nthis not /* this is a comment as well*/ this not /*/")
	lex := NewLexer(arr, len(arr), "../test/test.gosh")
	tokens, lerr := lex.Lex()

	if lerr != nil {
		log.Fatal(lerr.Error())
	}

	t.Log(TokenArrayToString(tokens))
}
