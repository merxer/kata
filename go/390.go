package main

import "fmt"

var (
	goIsFun bool = true
	name string = "pat"
)

func main() {
	const format = "%T(%v)\n"
	fmt.Printf(format, goIsFun, goIsFun)
	fmt.Printf(format, name, name)
}
