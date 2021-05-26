package main

import "fmt"
import "encoding/json"

func main() {
	a := struct {
		A int
	}{
		A: 10,
	}
	aBytes, err := json.Marshal(a)
	var bInt8 []int8
	for _, item := range aBytes {
		bInt8 = append(bInt8, int8(item))
	}
	fmt.Println(aBytes, err)
	fmt.Println(bInt8)
}
