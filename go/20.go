package main

import (
  "fmt"
  "time"
)

type Bootcamp struct {
  Lat float64
  Lon float64
  Date time.Time
}


func main() {
  fmt.Println(Bootcamp{
    Lat: 34.012836,
    Lon: -118.495538,
    Date: time.Now(),
  })
}
