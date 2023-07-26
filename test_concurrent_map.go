package main

import "fmt"
import "sync"

func main() {
	var mutex sync.Mutex
	m := map[int]int{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			m[i] = i + 1
		}(i)
	}
	wg.Wait()
	fmt.Println("m: ", m)
}
