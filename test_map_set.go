package main

import "fmt"

func main() {
	f()
	fmt.Println("vim-go")
}

var a map[uint64]uint64 = map[uint64]uint64{1: 1}

func f() {
	b := map[uint64]uint64{1: 1}
	a = b
}
