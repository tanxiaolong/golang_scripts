package main

import (
	"fmt"
	"sync"
)

func main(){
	o := &sync.Once{}
	wg := sync.WaitGroup{}
	for i := 0;i<10;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			o.Do(func() {
				fmt.Println("hello")
			})
		}()
	}
	wg.Done()
	c := make(chan int)
	<-c
}
