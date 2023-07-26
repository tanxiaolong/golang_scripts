package main

import "fmt"
import "sync"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	wg := &sync.WaitGroup{}
	for _, item := range a {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			fmt.Println(item)
		}(item)
	}
	wg.Wait()
}
