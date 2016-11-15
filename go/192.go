package main

import "fmt"

var		name	string
var		age		int
var		location	string


func main() {
	name, age, location = "Pat", 35, "CRMU"
	fmt.Println(name, age, location)
}
