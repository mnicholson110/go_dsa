package minheap

import (
  "golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
  Heap []T
  Length int
}

func New[T constraints.Ordered]() *MinHeap[T] {
  return &MinHeap[T]{Heap: make([]T, 0), Length: 0}
}

func (h *MinHeap[T]) Insert(value T) {
  h.Heap = append(h.Heap, value)
  h.heapifyUp(h.Length)
  h.Length++
}

func (h *MinHeap[T]) Delete() T {
  if h.Length == 0 {
    panic("Heap is empty")
  }
  
  out := h.Heap[0]

  if h.Length == 1 {
    h.Length--
    h.Heap = h.Heap[:0]
    return out
  }
  
  h.Length--
  h.Heap[0] = h.Heap[h.Length]
  h.heapifyDown(0)

  return out
}

func (h *MinHeap[T]) heapifyUp(idx int) {
  if (idx == 0) {
    return
  }

  parentIdx := h.parentIdx(idx)

  if (h.Heap[parentIdx] > h.Heap[idx]) {
    h.Heap[parentIdx], h.Heap[idx] = h.Heap[idx], h.Heap[parentIdx]
    h.heapifyUp(parentIdx)
  }
}

func (h *MinHeap[T]) heapifyDown(idx int) {
  leftChildIdx := h.leftChildIdx(idx)
  rightChildIdx := h.rightChildIdx(idx)
  
  if (idx >= h.Length || leftChildIdx >= h.Length) {
    return
  }
  
  if (h.Heap[leftChildIdx] > h.Heap[rightChildIdx] && h.Heap[idx] > h.Heap[rightChildIdx]) {
    h.Heap[idx], h.Heap[rightChildIdx] = h.Heap[rightChildIdx], h.Heap[idx]
    h.heapifyDown(rightChildIdx)
  } else if (h.Heap[idx] > h.Heap[leftChildIdx]) {
    h.Heap[idx], h.Heap[leftChildIdx] = h.Heap[leftChildIdx], h.Heap[idx]
    h.heapifyDown(leftChildIdx)
  }
}

func (h *MinHeap[T]) parentIdx(idx int) int {
  return (idx - 1) / 2
}

func (h *MinHeap[T]) leftChildIdx(idx int) int {
  return 2 * idx + 1
}

func (h *MinHeap[T]) rightChildIdx(idx int) int {
  return 2 * idx + 2
}


