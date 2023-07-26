package main

import (
	"fmt"
	"time"
)

/*
➜ golang_scripts git:(master) ✗ go run test_goroutine_3.go
outside: 0 1
outside: 1 2
outside: 2 3
1 2
2 3
2 3
➜ golang_scripts git:(master) ✗ go run test_goroutine_3.go
outside: 0 1
outside: 1 2
outside: 2 3
1 3
2 3
2 3
*/
func main() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		fmt.Println("outside:", i, v)
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
}
