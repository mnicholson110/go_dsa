package stack

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type Stack[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (s *Stack[T]) Len() int {
	return s.length
}

func (s *Stack[T]) Push(value T) {
	node := &Node[T]{value: value}
	if s.head == nil {
		s.head = node
		s.tail = node
	} else {
		s.head.prev = node
		node.next = s.head
		s.head = node
	}
	s.length++
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		panic("Stack is empty")
	}

	node := s.head
	s.head = s.head.next
	s.length--

	if s.head == nil {
		s.tail = nil
	} else {
		s.head.prev = nil
	}

	return node.value
}

func (s *Stack[T]) Peek() T {
	if s.head == nil {
		panic("Stack is empty")
	}
	return s.head.value
}
