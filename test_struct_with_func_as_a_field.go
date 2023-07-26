package main

import "fmt"

type a struct {
	step func() int
}

func main() {
	aa := &a{step: func() int {
		return 1
	}}
	fmt.Println(aa.step())
	aa = &a{}
	fmt.Println(aa)
	fmt.Println(aa.step())
}
