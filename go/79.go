package main

import "fmt"

func main() {
  var foots float64 = 0.0
  fmt.Print("Enter distance in foots ")
  fmt.Scanf("%f", &foots)

  meters := foots * 0.3048
  fmt.Printf("\n %f foots = %f meters", foots, meters)
}
