package main

func main() {
	action := func() string{
		return "in function"
	}

	println(action())
}
