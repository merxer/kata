package main

import "fmt"

const (
	ConstExample = "Canstant"
)

var (
	universityName = "CMRU"
	room = 264
)

type Student struct {
	Firstname, Lastname string
	universityName string
}

func (s *Student) print() {
	fmt.Println(s.Firstname,
				s.Lastname,
				s.universityName)
}

func main() {
	u := &Student{
		Firstname: "TS",
		Lastname: "S",
		universityName: universityName,
	}

	u.print()
}
