package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1}
	fmt.Println(getTagPIDs(a, 0, 4))
	fmt.Println(getTagPIDs(a, 4, 4))
	fmt.Println(getTagPIDs(a, 8, 4))
	fmt.Println(getTagPIDs(a, 12, 4))

}

func getTagPIDs(a []int, index, count int) []int {
	start := int(index)
	length := int(count)
	//if length >= len(a) {
	//	return a
	//	}
	if start+length > len(a) {
		length = len(a) - start
	}
	if start+length < start {
		return []int{}
	}
	return a[start : start+length]
}
