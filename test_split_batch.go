package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	list := []string{"1", "2", "3", "4", "5"}
	rlt := splitBatchPIDs(list)
	fmt.Println(rlt)
}

func splitBatchPIDs(pids []string) [][]string {
	var rlt [][]string
	batch := 2
	var split []string
	for i, pid := range pids {
		pid := pid
		split = append(split, pid)
		if (i+1)%batch == 0 {
			rlt = append(rlt, split)
			split = []string{}
		}
		if i == len(pids)-1 && len(split) <= batch && len(split) != 0 {
			rlt = append(rlt, split)
		}
	}
	return rlt
}
