package main

import "fmt"

type (
  Celsius float64
  Fahrenheit float64
)

const (
  AbsoluteZeroF Fahrenheit= -273.15
  FreezingC  Celsius = 0
  BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

func main() {
  fmt.Println(FToC(AbsoluteZeroF))

  fmt.Println(CToF(FreezingC))

  fmt.Println(CToF(BoilingC))
}
