package main

import "fmt"

func main() {
	var str string
	a := str == ""
	b := len(str) == 0
	fmt.Println(a, b)
}
