package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	A string
}

func main() {
	b := &A{}
	a := ""
	err := json.Unmarshal([]byte(a), &b)
	fmt.Println(err)
}
