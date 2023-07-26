package main

import (
	"fmt"
)

type Person struct {
	Money int
}

func main() {
	p := Person{Money: 100}
	go func() {
		p.Money += 1000
	}()
	fmt.Printf("Money: %d\n", p.Money)
}
