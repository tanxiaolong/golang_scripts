package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("c append a,b", c)
	a = append(a, b...)
	fmt.Println("a append a,b", a)
	f := []int{}
	a = append(f, a...)
	fmt.Println("a append f,a", a)

	d := [2]int{6, 7}
	size := len(d)
	e := make([]int, size)
	copy(e, d[:])
	fmt.Println("d[:]", d[:])
	fmt.Println("eee", e)
	a = append(a, e...)
	fmt.Println("aaa", a)

	for i := range a {
		a = append(a, i)
		fmt.Printf("%d, %p\n", len(a), a)
	}
}
