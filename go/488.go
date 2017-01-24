package main

import "fmt"

func main() {
	a := [2][3]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Printf("%q\n", a)
	fmt.Printf("a[1][2] = %d\n", a[1][2])
}
