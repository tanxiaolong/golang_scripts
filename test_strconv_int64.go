package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("vim-go")
	a := "0000"
	aInt64, err := strconv.ParseInt(a, 10, 64)
	fmt.Println(aInt64, err, uint8(0))
}
