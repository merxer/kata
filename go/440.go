package main

import "fmt"

func runBeforeDie() {
	println("defer()")
}

func runPanic() {
	panic("Crash...")
}

func main() {
	defer runBeforeDie()
	runPanic()
	fmt.Println("Hello, World")
}
