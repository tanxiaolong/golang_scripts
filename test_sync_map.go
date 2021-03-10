package main

import "fmt"
import "sync"

func main() {

	a := sync.Map{}
	vv, ok := a.Load("1")
	fmt.Println(vv, ok) //one true
	vv, ok = a.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false
	vv, ok = a.Load("1")
	fmt.Println(vv, ok) //one true

	fmt.Println("///////////////////////////")
	sf := &SafeMap{
		m:    map[int]int{},
		lock: sync.Mutex{},
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sf.lock.Lock()
			defer sf.lock.Unlock()
			sf.m[i] = i
		}(i)
	}
	wg.Wait()

	fmt.Println(sf)
}

type SafeMap struct {
	lock sync.Mutex
	m    map[int]int
}
