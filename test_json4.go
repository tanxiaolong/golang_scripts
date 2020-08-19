package main

import (
	"encoding/json"
	"fmt"
)

type JsonStruct struct {
	String *string        `json:"s"`
	Number *[]interface{} `json:"number"`
}

var rawJson = []byte(`{
"s":"We do not provide a number","number":["a",1,1.2]
}`)

func main() {
	var s *JsonStruct
	err := json.Unmarshal(rawJson, &s)
	if err != nil {
		panic(err)
	}

	if s.String == nil {
		panic("String is missing or null!")
	}

	if s.Number == nil {
		panic("Number is missing or null!")
	}

	fmt.Printf("String: %s  Number: %v\n", *s.String, *s.Number)
}
