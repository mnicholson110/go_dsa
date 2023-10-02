package maxheap

import (
	"testing"
)

func TestIntMaxHeap(t *testing.T) {
	heap := New[int]()

	if heap.Length != 0 {
		t.Errorf("Expected length to be 0, got %d", heap.Length)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if heap.Length != 8 {
		t.Errorf("Expected length to be 8, got %d", heap.Length)
	}

	if heap.Delete() != 420 {
		t.Errorf("Expected 1, got %d", heap.Delete())
	}

	if heap.Delete() != 69 {
		t.Errorf("Expected 3, got %d", heap.Delete())
	}

	if heap.Delete() != 8 {
		t.Errorf("Expected 4, got %d", heap.Delete())
	}

	if heap.Delete() != 7 {
		t.Errorf("Expected 5, got %d", heap.Delete())
	}

	if heap.Length != 4 {
		t.Errorf("Expected length to be 4, got %d", heap.Length)
	}

	if heap.Delete() != 5 {
		t.Errorf("Expected 7, got %d", heap.Delete())
	}

	if heap.Delete() != 4 {
		t.Errorf("Expected 8, got %d", heap.Delete())
	}

	if heap.Delete() != 3 {
		t.Errorf("Expected 69, got %d", heap.Delete())
	}

	if heap.Delete() != 1 {
		t.Errorf("Expected 420, got %d", heap.Delete())
	}

	if heap.Length != 0 {
		t.Errorf("Expected length to be 0, got %d", heap.Length)
	}
}
