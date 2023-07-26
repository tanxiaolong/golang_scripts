package main

import "fmt"
import "sync"

func main() {
	a := map[int]int{1: 1}
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	fmt.Printf("%p\n", a)
	b := map[string]string{}
	b[""] = ""
	fmt.Printf("%v\n", b)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			b["1"] = fmt.Sprintf("%d", i)
		}(i)
		go func() { fmt.Println(b["1"]) }()
	}
	wg.Wait()
	fmt.Println(b)
	c := a
	fmt.Println(a, c)
	delete(c, 1)
	fmt.Println(a,c)
	tt := new([]int)
	ttt := make([]int,5)
	fmt.Println(tt,ttt)
	tt = &ttt
	fmt.Println(tt,ttt)
	ttt[1]=1
	fmt.Println(tt,ttt)
}
