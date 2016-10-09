package main

import "fmt"

func f(i int) {
  fmt.Println(i)
}

func main() {
  for i := 0; i< 10; i++ {
    go f(i)
  }

  var input string
  fmt.Scanln(&input)
}
