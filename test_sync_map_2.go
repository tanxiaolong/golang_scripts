package main

import "fmt"
import "sync"

func main() {

	sm := tmp()
	fmt.Println(sm)
	id := []int{}
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		id = append(id, v.(int))
		if v.(int) == 2 {
			return false
		}
		return true
	}
	sm.Range(f)
	fmt.Println("id: ", id)
	fmt.Println("len id:", len(id))
}

func tmp() sync.Map {
	a := sync.Map{}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a.Store(i, i)
		}(i)
	}
	wg.Wait()
	return a
}
