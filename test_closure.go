package main

import "fmt"

func main() {

	f := filterPidFunc()
	var filtered []string
	tt := []string{"1", "2", "3", "4", "5", "2"}
	for i := range tt {
		filtered = f(tt[i])
	}
	fmt.Println(tt)
	fmt.Println(filtered)
}

// filterPidFunc 过滤重复pid闭包
func filterPidFunc() func(pid string) []string {
	tmpPidList := make([]string, 0)
	uniqueMap := make(map[string]bool)
	return func(pid string) []string {
		if !uniqueMap[pid] {
			uniqueMap[pid] = true
			tmpPidList = append(tmpPidList, pid)
		}
		return tmpPidList
	}
}
