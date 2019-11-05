package main

import "fmt"
import "time"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	go func() {
		time.Sleep(1)
		fmt.Println(3)
	}()
	fmt.Println(4)
	fmt.Println(5)
}
