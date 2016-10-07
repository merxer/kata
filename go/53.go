package main

import "fmt"

type One int

func (o One) string() string {
  s := "ONE"
  return s
}

func main() {
  var a One = 1
  fmt.Println(a.string())
 
}
