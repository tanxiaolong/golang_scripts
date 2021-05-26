package main

import "fmt"

func main() {
	a := map[int]int{1: 1}
	for i:=0;i<100;i++{
		a[i] = i
	}
	fmt.Printf("%p\n", a)
}
