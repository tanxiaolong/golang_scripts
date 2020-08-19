package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	for i, num := range nums {
		if i+1 < 1 {
			panic("asd")
		}
		fmt.Println(i, num)
	}
}
