package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{}
	//b := map[int]int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		//	b[i] = i
	}
	for i := 0; i < 100000; i++ {
		go func() {
			a[0] = i + 1
			//b[0] = i + 1
		}()
		go func() {
			fmt.Printf("a: %d, %p, %p\n", a[0], &a[0], &a)
			//b[0] = i + 2
			//	fmt.Printf("b: %d, %p\n", b[0], &b) // can't print address of &b[0]
		}()
	}
	time.Sleep(1 * time.Second)
}
