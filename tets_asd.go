package main

import "fmt"

var a = &[]int{1, 2, 3}
var b = []int{4, 5, 6}
var i int

func f() int {
	i = 1
	a = &[]int{7, 8, 9}
	b = []int{7, 8, 9}
	return 0
}

func main() {
	(*a)[i] = f()
	b[i] = f()
	fmt.Println(*a)
	fmt.Println(b)
}
