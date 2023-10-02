package queue

import (
	"testing"
)

func TestIntQueue(t *testing.T) {
	s := New[int]()

	s.Enqueue(5)
	s.Enqueue(7)
	s.Enqueue(9)

	if s.Dequeue() != 5 {
		t.Error("Expected 5")
	}

	if s.Length != 2 {
		t.Error("Expected length 2")
	}

	s.Enqueue(11)
	if s.Dequeue() != 7 {
		t.Error("Expected 7")
	}
	if s.Dequeue() != 9 {
		t.Error("Expected 9")
	}
	if s.Peek() != 11 {
		t.Error("Expected 11")
	}
	if s.Dequeue() != 11 {
		t.Error("Expected 11")
	}
	if s.Length != 0 {
		t.Error("Expected length 0")
	}
	s.Enqueue(69)
	if s.Peek() != 69 {
		t.Error("Expected 69")
	}
	if s.Length != 1 {
		t.Error("Expected length 1")
	}
}

func TestStringQueue(t *testing.T) {
	s := Queue[string]{}

	s.Enqueue("world")
	s.Enqueue("hello")
	s.Enqueue("hello")

	if s.Dequeue() != "world" {
		t.Error("Expected world")
	}

	if s.Length != 2 {
		t.Error("Expected length 2")
	}

	s.Enqueue("world")
	if s.Dequeue() != "hello" {
		t.Error("Expected hello")
	}
	if s.Dequeue() != "hello" {
		t.Error("Expected hello")
	}
	if s.Peek() != "world" {
		t.Error("Expected world")
	}
	if s.Dequeue() != "world" {
		t.Error("Expected world")
	}
	if s.Length != 0 {
		t.Error("Expected length 0")
	}
	s.Enqueue("dassaword")
	if s.Peek() != "dassaword" {
		t.Error("Expected dassaword")
	}
	if s.Length != 1 {
		t.Error("Expected length 1")
	}
}
