package main

import "fmt"

type a struct {
	msg string
}

func main() {
	s := fmt.Sprintf("%s-%s", "1", "2")
	fmt.Println(s)
	s = fmt.Sprintf("%s-%s", "1", "2", "3")
	fmt.Println(s)
	s = fmt.Sprintf("%s-%s", "1")
	fmt.Println(s)

	aa := &a{msg: "nihao"}
	s = fmt.Sprintf("%+v", aa)
	fmt.Println(s)
}
