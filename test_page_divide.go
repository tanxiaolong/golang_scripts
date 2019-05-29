package main

import "fmt"

func main() {
	a := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(pageDivide(1, 2, a))
	fmt.Println(pageDivide(2, 2, a))
	fmt.Println(pageDivide(3, 2, a))
	fmt.Println(pageDivide(4, 2, a))
	fmt.Println(pageDivide(5, 2, a))
}

func pageDivide(page, pageSize int, f []interface{}) []interface{} {
	length := len(f)
	if length < pageSize {
		return f
	}
	start := (page - 1) * pageSize
	if start > length {
		return []interface{}{}
	}
	end := start + pageSize
	if end > length {
		end = length
	}
	return f[start:end]
}
