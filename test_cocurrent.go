package main

import "fmt"
import "time"

func main() {
	aa := []int{5, 6, 7}
	for _, j := range aa {
		go bb(j)
	}
	time.Sleep(10 * time.Second)
}

func bb(i int) {
	fmt.Println(i)
}
