package main

import "fmt"

func jumpFloor(target int) int {
	first := 1
	second := 2
	result := 0
	if target == first || target == second {
		return target
	} else {
		for i := 3; i <= target; i++ {
			result = first + second
			first = second
			second = result
		}
	}
	return result
}

func main() {
	fmt.Println(jumpFloor(3))

}
