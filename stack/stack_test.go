package stack

import (
  "testing"
)

func TestIntStack(t *testing.T) {
  s := Stack[int]{}

  s.Push(5)
  s.Push(7)
  s.Push(9)

  if s.Pop() != 9 {
    t.Error("Expected 9")
  }

  if s.Length != 2 {
    t.Error("Expected length 2")
  }
  
  s.Push(11)
  if s.Pop() != 11 {
    t.Error("Expected 11")
  }
  if s.Pop() != 7 {
    t.Error("Expected 7")
  }
  if s.Peek() != 5 {
    t.Error("Expected 5")
  }
  if s.Pop() != 5 {
    t.Error("Expected 5")
  }
  if s.Length != 0 {
    t.Error("Expected length 0")
  }
  s.Push(69)
  if s.Peek() != 69 {
    t.Error("Expected 69")
  }
  if s.Length != 1 {
    t.Error("Expected length 1")
  }
}

func TestStringStack(t *testing.T) {
  s := Stack[string]{}

  s.Push("world")
  s.Push("hello")
  s.Push("hello")

  if s.Pop() != "hello" {
    t.Error("Expected hello")
  }

  if s.Length != 2 {
    t.Error("Expected length 2")
  }
  
  s.Push("world")
  if s.Pop() != "world" {
    t.Error("Expected world")
  }
  if s.Pop() != "hello" {
    t.Error("Expected hello")
  }
  if s.Peek() != "world" {
    t.Error("Expected world")
  }
  if s.Pop() != "world" {
    t.Error("Expected world")
  }
  if s.Length != 0 {
    t.Error("Expected length 0")
  }
  s.Push("dassaword")
  if s.Peek() != "dassaword" {
    t.Error("Expected dassaword")
  }
  if s.Length != 1 {
    t.Error("Expected length 1")
  }
}
