package main

import "fmt"

type Txl struct {
	A int
}

func main() {
	txl := &Txl{
		A: 1,
	}
	a := map[interface{}]int{
		"a": 1,
		2:   1,
		txl: 1,
	}
	fmt.Println(a)
}
