package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	aa := 1
	bb := 0

	if aa == 1 || bb == 0 {
		fmt.Println("aa")
	}
	if aa != 1 || bb == 0 {
		fmt.Println("bb")
	}
}
