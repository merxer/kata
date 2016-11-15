package main

import "fmt"

func main() {
	planet := struct {
		name     string
		diameter int
	}{"earth", 12742}

	fmt.Println("%T(%v)", planet)
}
