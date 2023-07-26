package main

import "fmt"

func main() {
	a := ""
	deli := ","
	fmt.Println(len(deli))
	fmt.Println(countGenericString(a, deli[0]))
}

func countGenericString(s string, c byte) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			n++
		}
	}
	return n
}
