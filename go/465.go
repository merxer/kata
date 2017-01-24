package main

import "fmt"

func greeter(first, second int) bool {
	if first > second {
		return true
	}
	return false
}

func main() {
	first, second := 2, 5
	fmt.Printf("Value first(%d) > second(%d) ? ", first, second)
	fmt.Println(greeter(first, second))
}
