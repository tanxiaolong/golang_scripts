package main

import "fmt"
import "encoding/json"
import "bytes"

type s struct {
	a string `json:"a"`
}

func main() {
	str := &s{
		a: "b\"c\"",
	}
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(str)
	fmt.Println(str)
	fmt.Println(bf.String(), err)
	strBytes, err := json.Marshal(str)
	fmt.Println(string(strBytes), err)
}
