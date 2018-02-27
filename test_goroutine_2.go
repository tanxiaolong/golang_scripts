package main

import "fmt"
import "sync"
import "runtime"
import "github.com/petermattis/goid"

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	//wg.Add(20)
	wg.Add(10000)
	//for i := 0; i < 10000; i++ {
	//	go func() {
	//		fmt.Printf("i:%d,id:%d\n", i,goid.Get())
	//		wg.Done()
	//	}()
	//}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Printf("i:%d,id:%d\n", i, goid.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
