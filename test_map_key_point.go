package main

import "fmt"

func main() {
	a := map[string]int{
		"a": 1,
	}
	fmt.Printf("map's address: %p\n", a)
	fmt.Printf("map's key's address: %d\n", a["a"])
}
