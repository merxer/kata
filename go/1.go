package main

var (
  name string
  age int
  location string
)

func main() {
  name, age, location = "Prince Oberyn", 32, "Dorne"

  println(name,
    age,
    location,
  )
}
