package main

import "fmt"
import "strconv"

func main() {
	str := "ab\"c"
	fmt.Println(str)
	fmt.Println(strconv.Quote(str))
	fmt.Println(strconv.Quote("ab\"c"))
	str = strconv.Quote(str)
	fmt.Println(str[1 : len(str)-1])
str = "今天是个|好日子"
str = strconv.Quote(str)
fmt.Println(str)
}
