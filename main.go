package main

import (
  "fmt"
  "linkedlist"
)

func main() {
  list := linkedlist.LinkedList[int]{}

  list.Prepend(5)
  list.Prepend(7)
  list.Prepend(9)

  fmt.Println(list.GetAt(0))
  fmt.Println(list.GetAt(2))

  list.Remove(9)
}
