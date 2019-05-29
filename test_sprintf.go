package main

import "fmt"

func main() {
	s := fmt.Sprintf("%s-%s", "1", "2")
	fmt.Println(s)
	s = fmt.Sprintf("%s-%s", "1", "2", "3")
	fmt.Println(s)
	s = fmt.Sprintf("%s-%s", "1")
	fmt.Println(s)
}
