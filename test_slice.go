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
	fmt.Println(slice[2:100])
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
