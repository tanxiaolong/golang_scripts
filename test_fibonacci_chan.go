package main

import "fmt"

func main() {
	c := make(chan int, 10)
	n := 10
	go fib(n, c)
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

func fib(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
