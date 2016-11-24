package main

import "fmt"

const Pi = 3.14
const (
	StatusOk = 200
	StatusCreated = 201
)

func main() {
	const Greeting = "สวัสดีครับ"

	fmt.Println(Pi)
	fmt.Println(StatusOk)
	fmt.Println(StatusCreated)
	fmt.Println(Greeting)
}
