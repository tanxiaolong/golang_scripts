package main

import (
	"fmt"
)

func generator(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func main() {
	var c chan int
	c = make(chan int)
	go generator(c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %d\n", <-c)
	}

	fmt.Println("You're boring; I'm leaving.")
	close(c)
}
