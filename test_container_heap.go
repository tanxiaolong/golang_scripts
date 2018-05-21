package main

import (
	"container/heap"
	"fmt"
	"sort"
	// "strconv"
)

type HeapInt []int

func (h HeapInt) Len() int {
	return len(h)
}

func (h HeapInt) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h HeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *HeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	slice := make(HeapInt, 0)
	heap.Init(&slice)
	for i := 0; i < 10; i++ {
		heap.Push(&slice, i*i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(slice[i])
	}
	bb := sort.SearchInts(slice, 23)
	fmt.Println(bb)
}
