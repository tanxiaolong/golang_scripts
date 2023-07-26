package main

import "fmt"

func main() {
	aim := 20
	arr := []int{5, 2, 3}
	dp := []int{}

	for i := range arr {
		dp = append(dp, GetCount(arr[i], aim))
	}

	fmt.Println(min(dp))
}

func min(arr []int) int {
	min := arr[0]
	for i := range arr {
		if min >= arr[i] {
			min = arr[i]
		}
	}
	return min
}
func GetCount(item, aim int) int {
	if aim < item {
		return 0
	}
	aim = aim - item
	return GetCount(item, aim) + 1
}
