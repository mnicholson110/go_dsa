package stack

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type Stack[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{Value: value}
	if s.Head == nil {
		s.Head = node
		s.Tail = node
	} else {
		s.Head.Prev = node
		node.Next = s.Head
		s.Head = node
	}
	s.Length++
}

func (s *Stack[T]) Pop() T {
	if s.Head == nil {
		panic("Stack is empty")
	}

	node := s.Head
	s.Head = s.Head.Next
	s.Length--

	if s.Head == nil {
		s.Tail = nil
	} else {
		s.Head.Prev = nil
	}

	return node.Value
}

func (s *Stack[T]) Peek() T {
	if s.Head == nil {
		panic("Stack is empty")
	}
	return s.Head.Value
}
