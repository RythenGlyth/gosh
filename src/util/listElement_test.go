package util

import (
	"container/list"
	"testing"
)

func TestBasicListGet(t *testing.T) {

	l := list.New()

	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	l.PushBack("e")
	l.PushBack("f")
	l.PushBack("g")
	l.PushBack("h")
	l.PushBack("i")
	l.PushBack("j")
	l.PushBack("k")

	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for i, s := range arr {
		isE := ListGet(l, i)
		if isE == nil {
			t.Errorf("Element @%d was reported nil", i)
		}

		isS, ok := isE.Value.(string)
		if !ok {
			t.Errorf("Element @%d could not be cast to string", i)
		}

		if isS != s {
			t.Errorf("Wrong element @%d returned: is %s, should be %s",
				i, isS, s)
		}
	}
}
