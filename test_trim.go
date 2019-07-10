package main

import (
	"fmt"
	"strings"
)

func main() {
	a := " a bc "
	fmt.Printf("$$$$%s$$$\n", a)
	fmt.Printf("$$$$%s$$$\n", strings.Trim(a, " "))
	fmt.Println("===========")
	a = " "
	fmt.Printf("$$$$%s$$$\n", a)
	fmt.Printf("$$$$%s$$$\n", strings.Trim(a, " "))
	fmt.Println(strings.Trim(a, " ") == "")

	fmt.Println(strings.TrimSuffix("abcd", "cd"))
}
