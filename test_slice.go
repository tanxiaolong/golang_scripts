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

}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
