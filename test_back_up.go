package main

import "fmt"

const max = 10

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := []int{4, 5, 6, 7, 8, 9, 10, 11}

	gap := max - len(a)
	fmt.Println(gap, len(b))
	a = append(a, b[:gap]...)
	fmt.Println(a, len(a), b[gap])
}
