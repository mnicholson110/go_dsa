package queue

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type Queue[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{Value: value}

	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		node.Prev = q.Tail
	}

	q.Tail = node
	q.Length++
}

func (q *Queue[T]) Dequeue() T {
	if q.Head == nil {
		panic("Queue is empty")
	}

	node := q.Head
	q.Head = q.Head.Next
	q.Length--

	if q.Head == nil {
		q.Tail = nil
	} else {
		q.Head.Prev = nil
	}

	return node.Value
}

func (q *Queue[T]) Peek() T {
	if q.Head == nil {
		panic("Queue is empty")
	}
	return q.Head.Value
}
