package main

import "fmt"

func main() {
	a := "宿迁市"
	l := len(a)
	fmt.Println(a, a[0:l-3])

	a = "1234"
	l = len(a)
	fmt.Println(a[0:3])
}
