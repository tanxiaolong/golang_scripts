package main

import "time"
import "fmt"
import "runtime"

func main(){
	ch := make(chan struct{})

	go func(ch chan struct{}){
		go func() {
			select {
				case <-ch:
					fmt.Println("interupt")
					runtime.Goexit() 
			}
		}()
		time.Sleep(10*time.Second)
		fmt.Println("end")
	}(ch)

	time.Sleep(3*time.Second)
	ch <- struct{}{}
	fmt.Println("main end")
}
