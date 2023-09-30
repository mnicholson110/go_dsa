package main

import (
  "linkedlist"
  "fmt"
)

func main() {
  intList := linkedlist.LinkedList[int]{}
  intList.Append(1)
  intList.Append(2)
  intList.Append(3)
  intList.Append(4)
  intList.Append(5)

  node := intList.Head
  for i := 0; i < intList.Length; i++ {
    fmt.Println(node.Value)
    node = node.Next
  }

  strList := linkedlist.LinkedList[string]{}
  strList.Append("a")
  strList.Append("b")
  strList.Append("c")
  strList.Append("d")
  strList.Append("e")

  strNode := strList.Head
  for i := 0; i < strList.Length; i++ {
    fmt.Println(strNode.Value)
    strNode = strNode.Next
  }
}
