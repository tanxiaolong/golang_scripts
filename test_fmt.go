package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := &User{
		Name: "tanxiaolong",
		Age:  68,
	}
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	t := "123"
	fmt.Printf("%T\n", t)
}
