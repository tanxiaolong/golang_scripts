package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		if a[i] == 5 {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	fmt.Println(a)
}
