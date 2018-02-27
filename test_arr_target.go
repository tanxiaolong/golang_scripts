package main

import (
	"fmt"
	_ "reflect"
)

func foo(arr []int, target int) int {
	m := map[int]int{}
	ts := 0
	for _, v := range arr {
		m[v] += 1
		ts += v
	}
	sum := 0
	for _, v := range arr {
		if _, exists := m[target-v]; exists {
			sum += target - v
		}
	}
	return ts - sum
}
func main() {
	fmt.Println(foo([]int{2, 1, 3, 5, 4, 6, 4}, 7))
}
