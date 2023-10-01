import (
  "errors"
)

ckage linkedlist

import (
  "testing"
)

func TestLinkedList(t *testing.T) {
  list := &LinkedList[int]{}
  list.Append(1)
  list.Append(2)
  list.Append(3)
  list.Append(4)
  list.Append(5) 
  list.Append(6)
  list.Append(7)
  list.Append(8)
  list.Append(9)
  list.Append(10)

  if list.Length != 10 {
    t.Errorf("Expected length 10, got %d", list.Length)
  }

  if list.Head.Value != 1 {
    t.Errorf("Expected head value 1, got %d", list.Head.Value)
  }

  if list.Tail.Value != 10 {
    t.Errorf("Expected tail value 10, got %d", list.Tail.Value)
  }

  list.Remove(5)

  if list.Length != 9 {
    t.Errorf("Expected length 9, got %d", list.Length)
  }

  list.Prepend(0)

  if list.Length != 10 {
    t.Errorf("Expected length 10, got %d", list.Length)
  }

  if list.Head.Value != 0 {
    t.Errorf("Expected head value 0, got %d", list.Head.Value)
  }
}
