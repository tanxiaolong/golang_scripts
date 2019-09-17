package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := []interface{}{}
	test = append(test, "a")
	test = append(test, 1)
	fmt.Println(test)
}
