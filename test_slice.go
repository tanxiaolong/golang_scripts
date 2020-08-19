package main

import (
	"fmt"
)

type A struct {
	A int
	B int
}

func main() {
	slice := []byte{1, 2, 3, 4}
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)

	gids := []string{}
	gids = append(gids, "")
	gids = append(gids, "")
	fmt.Println(len(gids))

	//fmt.Println(gids[3])
	//fmt.Println("12"[0:111])
	//fmt.Println(slice[2:100])

	fmt.Println("////////////////////")
	a := make([]int, 5)
	fmt.Printf("a's addr: %p\n", a)
	b := a[:4]
	fmt.Printf("b's addr: %p\n", &b)
	a[1] = 5
	fmt.Println(b)
	a = append(a, 1)
	fmt.Printf("a's addr: %p\n", &a)
	a[1] = 6
	fmt.Println(b)
	fmt.Println(a)
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
