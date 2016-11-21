package main

import "fmt"

type Age int

func (a Age) show() {
	fmt.Println(a)
}


func main() {
	var age Age = 20
	age.show()
}
