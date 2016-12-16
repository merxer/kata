package main

import "fmt"

func main() {
	vals := []int{4, 2, 6}
	for _, v := range vals {
		v--
	}
	fmt.Println(vals)
}
