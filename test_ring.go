package main

import "fmt"
import "container/ring"

func main() {
	fmt.Println("vim-go")
	r := ring.New(10)
	fmt.Printf("%+v\n", r)
}
