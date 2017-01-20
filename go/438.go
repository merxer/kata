package main

import "fmt"

type User struct {
	Firstname, Lastname string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

type Customer struct {
	Id int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))

	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))
}
