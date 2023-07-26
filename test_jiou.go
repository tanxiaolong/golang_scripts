package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i&1 == 0 {
			fmt.Println("偶数:", i)
		}
	}
}
