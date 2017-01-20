package main

import "fmt"

func runBeforeDie() {
	println("Keep log, before exit program")
}

func main() {
	defer runBeforeDie()
	fmt.Println("Hello World")
}
