package main

const (
	StarHyperGiant = iota
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
