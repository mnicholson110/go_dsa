package linkedlist

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
    l.Head = node
  } else {
    l.Head.Prev = node
    node.Next = l.Head
  }

  l.Head = node
  l.Length++
}

func (l *LinkedList[T]) Remove(value T) {
  if l.Head == nil {
    return
  }

  if l.Head.Value == value {
    l.Head = l.Head.Next
    l.Head.Prev = nil
    l.Length--
    return
  }

  node := l.Head.Next
  for node != nil {
    if node.Value == value {
      if node.Next == nil {
        l.Tail = node.Prev
        l.Tail.Next = nil
      } else {
        node.Prev.Next = node.Next
        node.Next.Prev = node.Prev
      }

      l.Length--
      return
    }

    node = node.Next
  }
}
