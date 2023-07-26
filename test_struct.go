package main

import "fmt"

type Animal struct {
	name string
	legs int
}

func main() {
	zoo := make([]Animal, 10)
	for i, animal := range zoo {
		fmt.Println("i:", i)
		animal.legs = i
	}
	fmt.Println(zoo)
}
