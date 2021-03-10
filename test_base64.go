package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	input := []byte("hello world")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)
	decodeBytes, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Println(string(decodeBytes))
}
