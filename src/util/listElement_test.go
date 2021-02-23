package util

import (
	"container/list"
	"testing"
)

func TestBasicListGet(t *testing.T) {
	l := list.New()
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for _, s := range arr {
		l.PushBack(s)
	}

	for i, s := range arr {
		isE := ListGet(l, i)
		if isE == nil {
			t.Errorf("Element @%d was reported nil", i)
			t.FailNow()
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

func TestRuneListToString(t *testing.T) {
	arr := []rune{'o', 'm', 'e', 'g', 'a', 'l', 'u', 'l'}
	l := list.New()

	for _, r := range arr {
		l.PushBack(r)
	}

	line := RuneListToString(l)

	if line != "omegalul" {
		t.Errorf("Wrong string returned: %s", line)
	}
}

func TestPositionInList(t *testing.T) {
	l := list.New()

	l.PushBack('q')
	q := l.Back()
	l.PushBack('u')
	l.PushBack('o')
	o := l.Back()
	l.PushBack('q')
	q2 := l.Back()
	l.PushBack('a')
	a := l.Back()

	pos := PositionInList(l, q)
	if pos != 0 {
		t.Errorf("Wrong position for element 0: %d", pos)
	}

	pos = PositionInList(l, a)
	if pos != 4 {
		t.Errorf("Wrong position for element 4: %d", pos)
	}

	pos = PositionInList(l, q2)
	if pos != 3 {
		t.Errorf("Wrong position for element 3: %d", pos)
	}

	pos = PositionInList(l, o)
	if pos != 2 {
		t.Errorf("Wrong position for element 2: %d", pos)
	}
}
