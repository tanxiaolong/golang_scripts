package main

import "fmt"
import "encoding/json"

func main() {
	fmt.Println("vim-go")
	str := []byte("abc")

	t := &[]byte{}
	json.Unmarshal(str, t)
	fmt.Println(t)

	//tmp := &string
	//Pass(tmp, str)
	//fmt.Println("tmp:", tmp)
}

func Pass(val interface{}, str []byte) {
	json.Unmarshal(str, val)
	//fmt.Println("pass:", *val)
}
