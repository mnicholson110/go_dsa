package main

import (
  "linkedlist"
  "fmt"
)

func main() {
  strList := linkedlist.LinkedList[string]{}
  strList.Append("this")
  strList.Append("is")
  strList.Append("a")
  strList.Append("linked")
  strList.Append("list")

  strNode := strList.Head
  for i := 0; i < strList.Length; i++ {
    fmt.Println(strNode.Value)
    strNode = strNode.Next
  }
}
