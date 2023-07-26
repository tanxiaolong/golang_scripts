package main

import "fmt"

func main() {
	fmt.Println(missed())
}

func missed() []int {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	length := len(a)
	missed := []int{}
	for i := 1; i <= length; i++ {
		for j := 0; j < length; j++ {
			if a[j]-i != 0 && j == length-1 {
				missed = append(missed, i)
			} else if a[j]-i == 0 {
				break
			}
		}
	}
	return missed
}
