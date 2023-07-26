package main

import "fmt"
import "math/rand"
import "time"

func main() {
	arr := []int{10, 12, 14, 16, 18}
	fmt.Println(arr)
	//fmt.Println(shuffle(arr))
	fmt.Println(shuffle2(arr))
	//r := randomRange(2,5)
	//fmt.Println(r)

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



func shuffle2(arr []int) []int{
	n :=len(arr)
	for i:=0;i<n;i++{
		r := randomRange(i+1,n-1)
		fmt.Println(r)
		arr[i],arr[r] = arr[r],arr[i]
	}
	return arr
}
func randomRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}