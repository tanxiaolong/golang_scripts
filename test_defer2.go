package main

import (
	"bytes"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var b *bytes.Buffer
	fmt.Println("Hello, playground", b.Bytes())
}
