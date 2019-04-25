package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(divide(1, 2, a))
	fmt.Println(divide(2, 2, a))
	fmt.Println(divide(3, 2, a))
	fmt.Println(divide(4, 2, a))
	fmt.Println(divide(5, 2, a))
}

func divide(page, pageSize int, i []int) []int {
	start := (page - 1) * pageSize
	end := start + pageSize
	return i[start:end]
}
