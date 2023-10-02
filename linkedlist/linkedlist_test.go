package linkedlist

import (
	"testing"
)

func TestIntLinkedList(t *testing.T) {
	list := New[int]()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if list.GetAt(2) != 9 {
		t.Errorf("Expected 9, got %d", list.GetAt(2))
	}

	list.RemoveAt(1)

	list.RemoveAt(0)

	list.RemoveAt(0)

	if list.Length != 0 {
		t.Errorf("Expected length of 0, got %d", list.Length)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if list.GetAt(0) != 9 {
		t.Errorf("Expected 9, got %d", list.GetAt(0))
	}

	if list.GetAt(2) != 5 {
		t.Errorf("Expected 5, got %d", list.GetAt(2))
	}

	list.Remove(9)

	if list.Length != 2 {
		t.Errorf("Expected length of 2, got %d", list.Length)
	}

	if list.GetAt(0) != 7 {
		t.Errorf("Expected 7, got %d", list.GetAt(0))
	}
}
