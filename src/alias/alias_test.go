package alias

import "testing"

func TestExpandSimply(t *testing.T) {
	m := Manager{}
	m.RegisterSimpleAlias("v", "vim")

	is, ok := m.expandSimply("v")
	if !ok {
		t.Errorf("Should have expanded")
	}
	compare(t, "expansion", is, "vim")

	is, ok = m.expandSimply("v foo")
	if !ok {
		t.Errorf("Should have expanded")
	}
	compare(t, "expansion", is, "vim foo")
}

func TestExpandGlobally(t *testing.T) {
	m := Manager{}
	m.RegisterGlobalAlias("F", "foo")

	is, ok := m.expandGlobally("F bar baz")
	if !ok {
		t.Errorf("Should have expanded")
	}
	compare(t, "expansion", is, "foo bar baz")

	is, ok = m.expandGlobally("bar F baz")
	if !ok {
		t.Errorf("Should have expanded")
	}
	compare(t, "expansion", is, "bar foo baz")

	is, ok = m.expandGlobally("bar baz F")
	if !ok {
		t.Errorf("Should have expanded")
	}
	compare(t, "expansion", is, "bar baz foo")
}

func TestMultiExpand(t *testing.T) {
	m := Manager{}
	m.RegisterSimpleAlias("gst", "g status")
	m.RegisterSimpleAlias("g", "git")
	m.RegisterGlobalAlias("C", "--colors=auto")

	compare(t, "multi expansion", m.Expand("gst C"), "git status --colors=auto")
}

func compare(t *testing.T, what, is, should string) {
	t.Helper()

	t.Logf("is: %s", is)
	t.Logf("should: %s", should)

	if is != should {
		t.Errorf("Wrong %s: got \"%s\", expected \"%s\"", what, is, should)
	}
}
