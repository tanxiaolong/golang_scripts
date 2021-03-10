package main

import (
	"fmt"
)

type A struct {
	A int
	B int
}

func main() {
	pageSize := 3
	pageNo := 10
	ian := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	start := pageNo * pageSize
	//end := pageSize * pageNo
	end := (pageNo + 1) * pageSize
	fmt.Println("start:", start, "end:", end)
	if end > len(ian) {
		end = len(ian)
	}
	if start >= len(ian) {
		fmt.Println("start > length")
	} else {
		fmt.Println(ian[start:end])
	}

	fmt.Println("////////////////////////////")
	var As []*A
	fmt.Println(As)
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

	////////////////////////
	fmt.Println("=====================")
	yy := []int{}
	fmt.Println(len(yy), cap(yy))
	yy = append(yy, 0)
	fmt.Println(len(yy), cap(yy))
	yy = append(yy, 1)
	fmt.Println(len(yy), cap(yy))
	yy = append(yy, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(len(yy), cap(yy))
	//yy = append(yy, 3)
	//fmt.Println(len(yy), cap(yy))
	//yy = append(yy, 4, 5, 6)
	//fmt.Println(len(yy), cap(yy))

	//yy = []int{1, 2}
	//fmt.Println(len(yy), cap(yy))
	//yy = append(yy, 3, 4, 5)
	//fmt.Println(len(yy), cap(yy))
	//xx := []byte{0, 1}
	//fmt.Println(len(xx), cap(xx))
	//xx = append(xx, 2, 3, 4)
	//fmt.Println(len(xx), cap(xx))
	fmt.Println("yyyyy", yy)
	fmt.Println("yyyyyyyyyyyyyyyyyyyyy: ", yy[:5])
	yy = yy[:5]
	fmt.Println("yyyyyyyyyyyyyyyyyyyyy: ", yy)
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
