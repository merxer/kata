package main

import "fmt"

func swap(a int, b int) (int, int) {
	return  b, a
}

func main() {
	a, b := 5, 2
	swapA, swapB := swap(a, b)
	fmt.Printf("%d %d", swapA, swapB)
}
