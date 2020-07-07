package main

import "fmt"
import "sync"
import "time"

func main() {
	end := 10000000
	t := time.Now()
	a(end)
	t2 := time.Now()
	fmt.Println(t2.Sub(t))

}

func a(end int) {
	wg := sync.WaitGroup{}
	c := make(chan int, end)
	for i := 0; i < end; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c <- i
		}(i)
	}
	wg.Wait()
	for len(c) != 0 {
		select {
		case rlt := <-c:
			{
				fmt.Println(rlt)
			}
		}
	}
}
