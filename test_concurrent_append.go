package main

import (
	"fmt"
	"sync"
)

func main() {
	var s []int
	var lock sync.Mutex
	wg := sync.WaitGroup{}
	for i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if i == 5 {
			s = append(s, i)
			continue
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			s = append(s, i)
		}(i)
	}
	wg.Wait()

	fmt.Println(s)
}
