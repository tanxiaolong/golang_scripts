package main

import (
	"fmt"
)

func main(){
	arr := []int{5,2,3,9,7,8,1}
	fmt.Println(arr)
	duipai(arr)
	fmt.Println(arr)
}
// 堆排序
func duipai(buf []int) {
    temp, n := 0, len(buf)

    for i := (n - 1) / 2; i >= 0; i-- {
        MinHeapFixdown(buf, i, n)
    }

    for i := n - 1; i > 0; i-- {
        temp = buf[0]
        buf[0] = buf[i]
        buf[i] = temp
        MinHeapFixdown(buf, 0, i)
    }
}

func MinHeapFixdown(a []int, i, n int) {
    j, temp := 2*i+1, 0
    for j < n {
        if j+1 < n && a[j+1] < a[j] {
            j++
        }

        if a[i] <= a[j] {
            break
        }

        temp = a[i]
        a[i] = a[j]
        a[j] = temp

        i = j
        j = 2*i + 1
    }
}
