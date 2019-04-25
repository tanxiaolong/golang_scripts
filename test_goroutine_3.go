package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
		}
	}()
	fmt.Println(runtime.NumGoroutine())
}
