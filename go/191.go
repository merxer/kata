package main

import "fmt"

var (
    name, location string
    age             int
)

func main(){
    name, location, age = "XYZ", "CRMU", 18
    fmt.Println(name, location, age)
}
