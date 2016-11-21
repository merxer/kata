package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Name string
	}
	Bootcamp struct {
		Event
		Lat, Lon float64
		Date time.Time
	}
)

func main() {
	event := &Bootcamp{}
	event.Name = "Go hackathon"
	event.Lat = 34.01
	event.Lon = -118.49
	event.Date = time.Now()

	fmt.Println(event)
}
