package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	rlt := []int{}
	tmp := []int{}

	for _, val := range test {
		tmp = []int{val}
		rlt = append(tmp, rlt...)
	}
	fmt.Println(rlt)
}
