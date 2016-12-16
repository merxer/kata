package main

import (
	"fmt"
)

const Max = 1000

func main() {
	work := make(chan int, Max)
	result := make(chan int)

	go func() {
		for i := 1; i< Max; i++ {
			if (i % 3) == 0 || (i % 5) == 0 {
				work <- i
			}
		}
		close(work)
	}()

	go func() {
		r := 0
		for i := range work {
			r = r + i
		}
		result <- r
	}()

	fmt.Println("Total:", <- result)
}
