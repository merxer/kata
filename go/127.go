package main

func square(x *float64) float64 {
  *x = *x * *x
  return *x
}

func main() {
  x := 1.5
  println(square(&x))
}
