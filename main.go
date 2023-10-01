package main

import (
  "fmt"
  "stack"
)

func main() {
  a := stack.Stack[int]{}

  a.Push(1)
  a.Push(2)
  a.Push(3)

  fmt.Println(a.Pop())
  fmt.Println(a.Pop())
  fmt.Println(a.Pop())

  b := stack.Stack[string]{}

  b.Push("this")
  b.Push("is")
  b.Push("a")
  b.Push("string")

  fmt.Println(b.Pop())
  fmt.Println(b.Pop())
  fmt.Println(b.Pop())
  fmt.Println(b.Pop())
  fmt.Println(b.Pop())  
}
