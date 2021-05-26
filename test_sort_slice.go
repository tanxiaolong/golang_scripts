package main

import "fmt"
import "sort"

func main() {
	var a []int
	a = []int{3, 2, 1, 5, 4, 8, 0}
	sort.Slice(a, func(i, j int) bool {
		if a[i] < a[j] {
			return true
		}
		return false
	})
	fmt.Println(a)
}
