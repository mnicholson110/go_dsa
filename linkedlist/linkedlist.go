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

func New[T comparable]() *LinkedList[T] {
  return &LinkedList[T]{
    Head: nil,
    Tail: nil,
    Length: 0,
  }
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
    panic("Index out of bounds")
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
      if current == l.Head {
        l.Head = l.Head.Next
        if l.Head != nil {
          l.Head.Prev = nil
        }
      } else if current == l.Tail {
        l.Tail = l.Tail.Prev
        l.Tail.Next = nil
      } else {
        current.Prev.Next = current.Next
        current.Next.Prev = current.Prev
      }

      l.Length--
      return
    }

    current = current.Next
  }
}

func (l *LinkedList[T]) RemoveAt(index int) {
  if index < 0 || index >= l.Length {
    panic("Index out of bounds")
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

func (l *LinkedList[T]) GetAt(index int) T {
  if index < 0 || index >= l.Length {
    panic("Index out of bounds")
  }

  current := l.Head
  for i := 0; i < index; i++ {
    current = current.Next
  }

  return current.Value}

