package main

import "fmt"

const (
	//A = 1
	A = 1 << iota
	B
	C
	D
	E
	F
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
