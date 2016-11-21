package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date time.Time
}

func main() {
	event := Bootcamp{
		34.01, -118.49, time.Now(),
	}

	fmt.Println(event)
}
