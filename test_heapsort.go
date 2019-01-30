package main

import (
	"fmt"
)

func main(){
	arr := []int{51,32,73,23,42,62,99,14,24,3943,58,65,80,120}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}
// 最小堆排序
func heapSort(input []int){
    inputLen := len(input)
    if inputLen == 0 {
        return
    }
    for i:=0; i<inputLen; i++{
        minAjust(input[i:])
    }
}
 
func minAjust(input []int){
    inputLen := len(input)
    if inputLen <= 1{
        return
    }
    for i:= inputLen/2 -1; i>=0; i--{
        if (2*i+1 <= inputLen-1) && (input[i] >= input[2*i+1]){
           input[i], input[2*i+1] = input[2*i+1], input[i]
        }
        if (2*i+2<= inputLen-1) && (input[i] >= input[2*i+2]){
            input[i], input[2*i+2] = input[2*i+2], input[i]
        }
    }
}

// 最大堆排序
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
