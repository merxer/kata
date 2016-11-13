package main

import "fmt"

var (
    name string
    age int
    location string
)
func main() {
    name, age, location = "AAA", 19, "CM"
    fmt.Println(name, age, location)
}
