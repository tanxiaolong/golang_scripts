//go 1.10
package main

import (
	"fmt"
)

// select sort
func ssort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := range arr {
		for j := range arr {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{3, 5, 2, 4, 1}
	fmt.Println(arr)
	ssort(arr)
	fmt.Println(arr)
}
