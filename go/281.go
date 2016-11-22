package main

import "fmt"

func runBeforeDIE() {
	println("Keep Log, Before Exit Program")
}

func runPanic() {
	panic("Crassh....")
}

func main() {
	defer runBeforeDIE()

	runPanic()
	fmt.Println("Hello World")
}

