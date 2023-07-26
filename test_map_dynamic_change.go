package main

import "fmt"

func main() {
	a := map[int]int{
		1: 1,
		2: 2,
		3: 1,
		4: 1,
		5: 1,
	}
	for p, i := range a {
		fmt.Println("index:", p)
		if a[i] == 1 {
			fmt.Println(len(a))
		} else {
			delete(a, i)
		}
	}
	fmt.Println("vim-go")
}
