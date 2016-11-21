package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date time.Time
}

func (b *Bootcamp) show() {
	fmt.Printf("event on %s, location (%f, %f)", b.Date, b.Lat, b.Lon)
}

func main() {
	event := &Bootcamp{
		34.01,
		-118.49,
		time.Now(),
	}

	event.show()
}
