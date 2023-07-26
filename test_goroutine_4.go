package main

import "fmt"
import "sync"

func main() {
	fmt.Println("vim-go")
	arr := []int{}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(tt int) {
			defer wg.Done()
			arr = append(arr, tt)
		}(i)
		fmt.Println(i)
	}
	wg.Wait()

	fmt.Println(arr)
}
