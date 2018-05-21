package main

import "fmt"
import "sync"

//import "runtime"
//import "time"

func main() {
	//runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("a: ", i)
			wg.Done()
		}()
		// without sleep, the loop goes so fast that the goroutine outputs 10 with no time to react..
		//time.Sleep(time.Second)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("b: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
