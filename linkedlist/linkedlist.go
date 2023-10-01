package linkedlist

import (
  "errors"
)

type Node[T comparable] struct {
  Value T
  Next *Node[T]
  Prev *Node[T]
}

type LinkedList[T comparable] struct {
  Head *Node[T]
  Tail *Node[T]
  Length int 
}

func (l *LinkedList[T]) Append(value T) {
  node := &Node[T]{Value: value}

  if l.Head == nil {
    l.Head = node
  } else {
    l.Tail.Next = node
    node.Prev = l.Tail
  }

  l.Tail = node
  l.Length++
}

func (l *LinkedList[T]) Prepend(value T) {
  node := &Node[T]{Value: value}

  if l.Head == nil {
    l.Tail = node
  } else {
    l.Head.Prev = node
    node.Next = l.Head
  }

  l.Head = node
  l.Length++
}

func (l *LinkedList[T]) AddAt(index int, value T) {
  if index < 0 || index > l.Length {
    panic(errors.New("Index out of bounds"))
  }

  if index == l.Length {
    l.Append(value)
    return
  }

  if index == 0 {
    l.Prepend(value)
    return
  }

  node := &Node[T]{Value: value}
  current := l.Head
  for i := 0; i < index; i++ {
    current = current.Next
  }

  node.Next = current
  node.Prev = current.Prev
  current.Prev.Next = node
  current.Prev = node
  l.Length++
}

func (l *LinkedList[T]) Remove(value T) {
  current := l.Head
  for current != nil {
    if current.Value == value {
      if current.Prev == nil {
        l.Head = current.Next
        l.Head.Prev = nil
      } else {
        tmp := current.Prev
        tmp.Next = current.Next
        tmp.Next.Prev = tmp
      }

      if current.Next == nil {
        l.Tail = current.Prev
        l.Tail.Next = nil
      } else {
        tmp := current.Next
        tmp.Prev = current.Prev
        tmp.Prev.Next = tmp 
      }

      l.Length--
      return
    }

    current = current.Next
  }
}

func (l *LinkedList[T]) RemoveAt(index int) {
  if index < 0 || index >= l.Length {
    panic(errors.New("Index out of bounds"))
  }

  if index == 0 {
    l.Head = l.Head.Next
    if l.Head != nil {
      l.Head.Prev = nil
    }
  } else if index == l.Length - 1 {
    l.Tail = l.Tail.Prev
    l.Tail.Next = nil
  } else {
    current := l.Head
    for i := 0; i < index; i++ {
      current = current.Next
    }

    current.Prev.Next = current.Next
    current.Next.Prev = current.Prev
  }

  l.Length--
}

func (l *LinkedList[T]) GetAt(index int) (value T, err error) {
  if index < 0 || index >= l.Length {
    return value, errors.New("Index out of bounds")
  }

  current := l.Head
  for i := 0; i < index; i++ {
    current = current.Next
  }

  return current.Value, nil
}

