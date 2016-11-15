package main

const (
	StarHyperGiant = 2.0 * iota
	StarSuperGiant
	StarBrightGiant
	StarGiant
	StarSubGiant
	StarDwarf
	StarSubDwarf
	StarWhiteDwarf
	StarRedDwarf
	StarBrownDwarf
)

func main() {
	println(StarHyperGiant,
		StarSuperGiant,
		StarBrightGiant,
		StarGiant,
		StarSubGiant,
	)
}
