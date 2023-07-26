package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(5 * time.Second)
	//go func() {
	//for now := range c {
	//	fmt.Println(now)
	//	fmt.Println("5s later")
	//}

	for range c {
		fmt.Println("5s lllater")
	}
	//}()
}
