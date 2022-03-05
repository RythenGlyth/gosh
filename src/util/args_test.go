package util_test

import (
	"gosh/src/util"
	"reflect"
	"testing"
)

func TestArgs(t *testing.T) {
	line := "kill 1 -1 -hsq 5 3 --verbose --timeout 15 16 -t 17 foo asdf -- myprog \"sp c\" s\\ pc --timeout"

	sparams := map[rune]int{
		's': 1,
		'q': 1,
		't': 2,
	}

	lparams := map[string]int{
		"timeout": 2,
	}

	is, err := util.ParseLine(line, sparams, lparams)
	if err != nil {
		t.Error(err)
	}

	should := &util.Args{
		Prog:    "kill",
		UParams: []string{"1", "asdf", "myprog", "sp c", "s pc", "--timeout"},
		NParams: map[string][]string{
			"s":       {"5"},
			"q":       {"3"},
			"timeout": {"15", "16"},
			"t":       {"17", "foo"},
		},
		SOps: []rune{'1', 'h'},
		LOps: []string{"verbose"},
	}

	if !reflect.DeepEqual(is, should) {
		t.Error("is != should")
	}

	cmpargs(t, is, should)
}

func cmpargs(t *testing.T, is, should *util.Args) {
	t.Helper()

	t.Log("is:\n", is)
	t.Log("should:\n", should)

	if is.Prog != should.Prog {
		t.Errorf("wrong prog, got %s, should be %s", is, should)
	}

	cmpsarr(t, "uparams", is.UParams, should.UParams)
	cmpmap(t, "nparams", is.NParams, should.NParams)
	cmprarr(t, "sops", is.SOps, should.SOps)
	cmpsarr(t, "lops", is.LOps, should.LOps)
}

func cmpsarr(t *testing.T, what string, is, should []string) {
	t.Helper()

	if len(is) != len(should) {
		t.Errorf("wrong %s length, got %d, expected %d",
			what, len(is), len(should))
		return
	}

	for i, iss := range is {
		if should[i] != iss {
			t.Errorf("wrong %s element at %d, got %s, expected %s",
				what, i, iss, should[i])
		}
	}
}

func cmprarr(t *testing.T, what string, is, should []rune) {
	t.Helper()

	if len(is) != len(should) {
		t.Errorf("wrong %s length, got %d, expected %d",
			what, len(is), len(should))
		return
	}

	for i, iss := range is {
		if should[i] != iss {
			t.Errorf("wrong %s element at %d, got %q expected %q",
				what, i, iss, should[i])
		}
	}
}

func cmpmap(t *testing.T, what string, is, should map[string][]string) {
	t.Helper()

	if len(is) != len(should) {
		t.Errorf("wrong %s length, got %d, expected %d",
			what, len(is), len(should))
		return
	}

	for idx, valS := range should {
		valI, ok := is[idx]
		if !ok {
			t.Errorf("missing index %s", idx)
		}

		if !reflect.DeepEqual(valS, valI) {
			t.Errorf("value differs for %s", idx)
			t.Logf("got %v", valI)
			t.Logf("got %v", valS)
		}
	}
}
