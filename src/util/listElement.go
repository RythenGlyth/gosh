package util

import (
	"container/list"
	"strings"
)

// ListGet returns the index-th element from the specified list.
// If the list does not contain enough elements, nil is returned.
func ListGet(l *list.List, index int) *list.Element {
	if l.Len() <= index {
		return nil
	}

	var e *list.Element

	if index <= l.Len()/2 {
		e = l.Front()

		for i := 0; i < index; i++ {
			e = e.Next()
		}
	} else {
		e = l.Back()

		for i := 1; i < l.Len()-index; i++ {
			e = e.Prev()
		}
	}

	return e
}

// RuneListToString creates a string from a linked list of runes.
// If any of the elements is not a rune, the function panics.
// The function takes O(n).
func RuneListToString(l *list.List) string {
	if l == nil || l.Front() == nil {
		return ""
	}

	var line strings.Builder

	e := l.Front()

	for e != l.Back() {
		line.WriteRune(e.Value.(rune))
		e = e.Next()
	}
	line.WriteRune(e.Value.(rune))

	return line.String()
}

// PositionInList returns the position of an element in the list.
// If the element was not found, -1 is returned.
func PositionInList(l *list.List, elem *list.Element) int {
	if l == nil || l.Front() == nil || elem == nil {
		return -1
	}

	var i int

	e := l.Front()

	for e != l.Back() {
		if e == elem {
			return i
		}

		i++

		e = e.Next()
	}

	if e == elem {
		return i
	}

	return -1
}
