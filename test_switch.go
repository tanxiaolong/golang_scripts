package main

import "fmt"

func main() {
	a := "a"
	b := "b"
	switch a {
	case "a":
		if b == "b" {
			fmt.Println("555")
			break
		}
		fmt.Println("666")
	case "b":
		fmt.Println("777")
	default:
		fmt.Println("888")
	}
}
