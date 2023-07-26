package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	BB *B `json:"b"`
}

type B struct {
	C int `json:"c"`
}

func main() {
	//aa := &A{}
	var aa A
	bb := &B{}
	//err := json.Unmarshal([]byte(`{"b":{"c":1}}`), aa)
	err := json.Unmarshal([]byte(`{"c":1}`), bb)
	aa.BB = bb
	fmt.Println(aa, aa.BB, bb, err)
	//fmt.Println(aa, aa.BB, aa.BB.C, err)
}
