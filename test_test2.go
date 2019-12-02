package main

import "fmt"

func main() {
	prefix := "v"
	fmt.Println("vim-go"[len(prefix):])
}
