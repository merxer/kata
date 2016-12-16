package main

import "fmt"

func main() {
	for i, v := range "Hello" {
		fmt.Println(i, v)
	}
}
