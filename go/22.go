package main

import "time"

func main() {
  timer := time.NewTimer(time.Second)

  go func() {
    <- timer.C
    println("Timer expired")
  }()

  stop := timer.Stop()
  println("timer cancelled: ", stop)
}
