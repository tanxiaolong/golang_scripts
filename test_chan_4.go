package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {

		for i := 0; i < 100000; i++ {
			if i%2 == 0 {
				c1 <- i
			} else {
				c2 <- i
			}
		}
	}()

	for i := 0; i < 100000; i++ {
		select {

		case i := <-c1:
			go fmt.Println(i, ":c1")
		case i := <-c2:
			go fmt.Println(i, ":c2")

		}

	}

}
