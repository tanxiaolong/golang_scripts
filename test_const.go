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

const i = 100
var j = 123

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	
	fmt.Println(&j,j)
	fmt.Println(&i,i)
}
