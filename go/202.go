package main

import "math"

const i = "G is " + " for Go"
const j = 'V'
const k1, k2 = true, !k1
const m1 = math.Pi / 3.141592

func main() {
	println(i,
		k1,
		k2,
		m1,
	)
}
