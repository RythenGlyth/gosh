package builtins_test

import (
	"gosh/src/alias"
	"gosh/src/builtins"
	"gosh/src/shared"
	"gosh/src/test_util"
	"testing"
)

type AliasStub struct {
	test_util.GoshStub
	asm *alias.Manager
}

func (f AliasStub) GetAliasManager() shared.AliasManager {
	return f.asm
}

func TestAlias(t *testing.T) {
	g := &AliasStub{test_util.GoshStub{T: t}, &alias.Manager{}}
	a := &builtins.Alias{}

	err := a.Eval(g, "alias foo=bar")
	if err != nil {
		t.Error(err)
	}

	err = a.Eval(g, "alias me you")
	if err != nil {
		t.Error(err)
	}

	err = a.Eval(g, "alias -g F=--force")
	if err != nil {
		t.Error(err)
	}

	t.Log("Simple Aliases:")
	is := g.GetAliasManager().ListSimpleAliases()
	for _, a := range is {
		t.Log(" -", a.Format())
	}

	if len(is) != 2 {
		t.Errorf("Wrong amount of simple aliases, got %d, expected 2", len(is))
		return
	}

	txt := is[0].Format() + ", " + is[1].Format()
	should := "foo aliased to bar, me aliased to you"
	if txt != should {
		t.Errorf("Wrong aliases")
		t.Logf("Got %s", txt)
		t.Logf("Expected %s", should)
	}

	t.Log("Global Aliases:")
	is = g.GetAliasManager().ListGlobalAliases()
	for _, a := range is {
		t.Log(" -", a.Format())
	}

	if len(is) != 1 {
		t.Errorf("Wrong amount of global aliases, got %d, expected 1", len(is))
	}

	txt = is[0].Format()
	should = "F aliased to --force"
	if txt != should {
		t.Errorf("Wrong aliases")
		t.Logf("Got %s", txt)
		t.Logf("Expected %s", should)
	}
}
