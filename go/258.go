package main

import "fmt"

var pow = []int{1, 2, 8, 16, 32, 64, 128}

func main() {
	for _, value := range pow {
		if value > 16 {
			break
		}
		fmt.Printf("%d\n", value)
	}
}
