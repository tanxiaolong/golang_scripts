package main

import "fmt"

func main() {
	fmt.Println(getSum2(-1, 1))
}
func getSum2(a, b int) int {
	return a ^ b
}
func getSum(a int, b int) int {
	sum := 0
	if a > 0 {
		for i := 0; i < a; i++ {
			sum++
		}
	} else {
		for i := a; i < 0; i++ {
			sum--
		}
	}
	if b > 0 {
		for i := 0; i < b; i++ {
			sum++
		}
	} else {
		for i := b; i < 0; i++ {
			sum--
		}
	}
	return sum
}
