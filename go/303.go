package main

import "fmt"

func swap(a, b int) (int, int) {
	x, y := b, a
	return x, y
}

func main() {
	a, b := 5, 2
	swapA, swapB := swap(a, b)
	fmt.Printf("swap a(%d), b(%d) is %d, %d\n", a, b, swapA, swapB)
}
