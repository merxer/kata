package main

import "fmt"

func main() {
	cities := make([]string, 2)

	cities[0] = "TAk"
	cities[1] = "Buriram"
	cities = append(cities, "San Diego", "BAngkok")

	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))

	cities = append(cities, "Surin")
	fmt.Printf("%q\n", cities)
	fmt.Println(len(cities))
}
