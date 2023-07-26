package main

import "runtime"
import "sync"
import "fmt"

func main(){
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i:=1;i<=258;i++{
		wg.Add(1)
		go func(n int){
			defer wg.Done()
			fmt.Printf("%d ",n)
		}(i)
	}
	wg.Wait()
}
