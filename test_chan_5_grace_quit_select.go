package main

import (
	"fmt"
)

func main() {
	bbb := map[int][]int{1: {1, 2, 3}, 2: {2, 3, 4}, 3: {3, 4, 5}, 5: {5, 6, 7}}
	aaa := []int{1, 2, 3, 4, 5}

	c1 := make(chan int)
	stopChan := make(chan int)
	go func() {
		for i := range aaa {
			item := aaa[i]
			makeData(c1, bbb[item])
		}
		stopChan<-1
	}()
	for {
		select {
		case <-stopChan:
			goto end
		case val := <-c1:
			fmtp(val)

		}
	}
	end:
}

func makeData(aaa chan int, ddd []int) {
	for i := range ddd {
		val := ddd[i]
		aaa <- val * val
	}
}

func fmtp(val int) {
	fmt.Println(val)
}
