package main

import "fmt"

func main(){
  x := make(map[int]int)
  x[1] = 10

  fmt.Println(x[1])
  fmt.Println(len(x))

  delete(x, 1)
  fmt.Println(x)
  fmt.Println(len(x))
}
