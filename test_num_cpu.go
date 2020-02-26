package main

import (
	"fmt"
	"runtime"
)

func main(){

	fmt.Println("num of cpus: ",runtime.NumCPU())
	fmt.Println("num of goroutine: ",runtime.NumGoroutine())

	
}