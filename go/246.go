package main

import "fmt"

func main() {
	a := [2][3]string{
		{"11", "12", "13"},
		{"21", "22", "23"},
	}

	fmt.Printf("%q\n", a)
}
