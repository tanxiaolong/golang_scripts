package main

import "fmt"

func main() {
	fmt.Println(test(8, 2))
}
func test(n, m int) int {
	round := n / m
	sum := 0
	turn := false
	for i := 0; i < round; i++ {
		fmt.Println("turn:", turn)
		if !turn {
			for j := i * m; j < m; j++ {
				sum = sum - (j + 1)
				fmt.Println("turn:", turn, sum)
				turn = true
			}
		}
		if turn {
			for j := i * m; j < m; j++ {
				sum = sum + (j + 1)
				fmt.Println("turn:ture", sum)
				turn = false
			}
		}

	}
	return sum
}
