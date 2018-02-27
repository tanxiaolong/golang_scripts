package main

import "fmt"
import "math/rand"
import "time"

func main() {
	arr := []int{10, 12, 14, 16, 18}
	fmt.Println(shuffle(arr))
}

func shuffle(arr []int) []int {
	rlt := make([]int, len(arr))
	// 时间种子
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(arr)) {
		rlt[i] = arr[j]
	}
	return rlt
}
