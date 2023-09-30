package linkedlist

type Node struct {
  Value int
  Next *Node
  Prev *Node
}

type LinkedList struct {
  Head *Node
  Tail *Node
  Length int
}

func (l *LinkedList) Append(value int) {
  node := &Node{Value: value}

  if l.Head == nil {
    l.Head = node
  } else {
    l.Tail.Next = node
    node.Prev = l.Tail
  }

  l.Tail = node
  l.Length++
}

func (l *LinkedList) Prepend(value int) {
  node := &Node{Value: value}

  if l.Head == nil {
    l.Head = node
  } else {
    l.Head.Prev = node
    node.Next = l.Head
  }

  l.Head = node
  l.Length++
}

func (l *LinkedList) Remove(value int) {
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
