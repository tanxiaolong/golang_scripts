package main

import "fmt"

// The range expression, a, is evaluated ONCE before beginning the loop.
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
