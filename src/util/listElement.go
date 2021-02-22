package util

import "container/list"

// ListGet returns the index-th element from the specified list.
// If the list does not contain enough elements, nil is returned
func ListGet(l *list.List, index int) *list.Element {
	if l.Len() <= index {
		return nil
	}

	var e *list.Element

	if index <= l.Len()/2 {
		e = l.Front()
		var i int
		for i = 0; i < index; i++ {
			e = e.Next()
		}
	} else {
		e = l.Back()
		var i int
		for i = 1; i < l.Len()-index; i++ {
			e = e.Prev()
		}
	}

	return e
}
