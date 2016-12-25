package main

import "fmt"

func main() {
	for {
		i := 1
		fmt.Printf("%d kata\n", &i)
		i = i + 1
	}
}
