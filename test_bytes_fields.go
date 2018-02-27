//go 1.10
package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := []byte("hello world")
	fields := bytes.Fields(str)
	fmt.Printf("%+v\n", fields)
	fmt.Printf("filed0 cap: %d. field1 cap: %d. \n", len(fields[0]), len(fields[1]))
	_ = append(fields[0], 'X', 'Y')
	fmt.Println(string(fields[0]))
	fmt.Printf("old string: %s. field0: %s. field1: %s. \n", string(str), fields[0], fields[1])

}
