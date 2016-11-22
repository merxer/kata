package main

import "fmt"

type myStr string

func (m myStr) show() {
	fmt.Println(m)
}

func main() {
	var name myStr = "Pat"
	fmt.Println(name)

	name.show()
}

