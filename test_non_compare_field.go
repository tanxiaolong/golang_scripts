package main

import "fmt"

type A struct {
	a []string
	B int
}

type B struct {
	a map[int]int
}

func main() {
	var a, b A
	_ = a == b
	fmt.Println("vim-go")

	// slice
	var c, d []string
	_ = c == d

	// chan
	var e, f chan int
	_ = e == f

	// map
	var g, h map[int]int
	_ = g == h

	// func
	var i, j func()
	_ = i == j

	var m, n B
	_ = m == n
}
