package main

import "fmt"
import "sync"

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var a, b int
	go return2(wg, &a)
	go return3(wg, &b)
	wg.Wait()

	var c = a + b
	fmt.Println("a , b, c: ", a, b, c)
}

func return2(wg *sync.WaitGroup, a *int) {
	defer wg.Done()
	*a = 2
	return
}

func return3(wg *sync.WaitGroup, b *int) {
	defer wg.Done()
	*b = 3
	return
}
