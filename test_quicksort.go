// go 1.10
package main

import "fmt"

func main() {
	arr := []int{3, 6, 1, 7, 9, 4, 5}
	//arr := []int{1,3,4,5,6,7,9}
	//arr := []int{9,7,6,5,4,3,1}
	fmt.Println(arr)
	qsort(arr)
	fmt.Println(arr)
}

func qsort(a []int) {
	if len(a) < 2 {
		return
	}
	l, r := 0, len(a)-1
	// 可以没有
	//mid := len(a) / 2
	//a[r], a[mid] = a[mid], a[r]
	for i := range a {
		if a[i] < a[r] {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}
	a[r], a[l] = a[l], a[r]
	qsort(a[:l])
	qsort(a[l+1:])
}
