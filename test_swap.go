package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
func main() {
	a := 1
	b := 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}
