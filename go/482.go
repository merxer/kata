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
	eventName := &Event{Name: "Go"}

	event := &Bootcamp{
		Event: *eventName,
		Lat: 34.012,
		Lon: -118.495,
		Date: time.Now(),
	}

	fmt.Println("%v", event)
}
