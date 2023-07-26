package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/*
   (浮点数, 需要保留的小数位数)
   不会四舍五入，如果位数不够不会补0
*/
func float64Save(data float64, n int) float64 {
	temp := fmt.Sprintf("%v", data)
	list := strings.Split(temp, ".")
	if len(list) < 2 {
		return data
	}
	length := len(list[1])
	if length <= n {
		return data
	} else {
		list[1] = string([]byte(list[1])[:n])
		theString := strings.Join(list, ".")
		theFloat64, _ := strconv.ParseFloat(theString, 64)
		return theFloat64
	}
}
func main() {
	fmt.Println(float64Save(0, 4))
	fmt.Println(float64Save(100, 4))
	fmt.Println(float64Save(-100, 4))
	fmt.Println(float64Save(100.12345, 4))
	fmt.Println(float64Save(-100.12345, 4))

	a := 12.3
	b := fmt.Sprintf("%.2f", a)
	fmt.Println(b)

	aa := "{\"a\":12.3}"
	test := &A{}
	json.Unmarshal([]byte(aa), &test)
	fmt.Println(test)
}

type A struct {
	A int `json:"a"`
}
