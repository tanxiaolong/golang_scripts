//go 1.10
package main

import (
	"fmt"
	"sync"
)

func chanCounter() chan int {
	c := make(chan int)
	go func() {
		for x := 1; ; x++ {
			c <- x
		}
	}()
	return c
}

func mutexCounter() func() int {
	var m sync.Mutex
	var x int

	return func() (n int) {
		m.Lock()
		x++
		n = x
		m.Unlock()

		return
	}
}
func main() {
	c := chanCounter()
	fmt.Println(<-c)
	m := mutexCounter()
	fmt.Println(m())
}
