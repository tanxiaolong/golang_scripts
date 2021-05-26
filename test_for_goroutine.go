package main

import "fmt"
import "time"

func main() {
	pids := []string{"1", "2", "3", "4"}
	for i, val := range pids {
		i, val := i, val
		go func() {
			fmt.Println(i, val)
		}()
		//go aaa(val)
	}
	time.Sleep(1 * time.Second)
}

func aaa(i string) {
	fmt.Println(i)
}
