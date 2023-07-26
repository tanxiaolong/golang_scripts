package main

import "sync"
import "fmt"

func main(){
  s := []int{}
  wg := sync.WaitGroup{}
  l := sync.Mutex{}
  for i:=0;i<10;i++ {
	i := i
	wg.Add(1)
	go func(in int){
		l.Lock()
		defer l.Unlock()
		s = append(s, in)
		fmt.Println(in, s)
		wg.Done()
	}(i)
  }
  wg.Wait()
  fmt.Println(s)
}
