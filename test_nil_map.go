package main

import "fmt"

func main() {
	a := map[string]int{}
	a = nil
	if _, exists := a["time"]; exists {
		fmt.Println(321)
	}
	fmt.Println("vim-go")
	m := nilMap()
	val, ok := m[1]
	fmt.Println(val, ok)
}

func nilMap() map[int]int {
	return nil
}
