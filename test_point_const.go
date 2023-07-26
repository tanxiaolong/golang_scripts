package main

import "fmt"

type A struct {
	a *uint32
}

func main() {
	aa := &A{
		a: &(uint32)(1),
	}
	fmt.Println("vim-go", aa)
}
