package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	list = []int{1}
	list = []int{1, 2, 3, 4}
	list = []int{1, 2, 3, 4, 5}
	list = []int{}
	list = []int{1, 2, 3, 4, 5, 6}
	list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rlt := splitBatchPIDs(list)
	fmt.Println(rlt)
}

func splitBatchPIDs(list []int) [][]int {
	var rlt [][]int
	batch := 5
	var split []int
	for i := range list {
		split = append(split, list[i])
		if (i+1)%batch == 0 {
			rlt = append(rlt, split)
			split = []int{}
		}
		if i == len(list)-1 && len(split) <= batch && len(split) != 0 {
			rlt = append(rlt, split)
		}
	}
	return rlt
}
