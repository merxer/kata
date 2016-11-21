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
	eventName := Event{Name: "Go Hackathon"}

	event := &Bootcamp{
		eventName,
		34.01,
		-118.49,
		time.Now(),
	}

	fmt.Println(event)
}
