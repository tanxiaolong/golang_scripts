package main

import "fmt"

func main() {
	for i := 0; i < 100000; i++ {
		if i < 1000 {
			i = i + 2
			continue
		}
		go func() {
			fmt.Println(i)
		}()
	}
}
