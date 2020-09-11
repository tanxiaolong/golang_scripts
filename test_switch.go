package main

import "fmt"

func main() {
	a := "b"
	b := "b"
	aa := "aa"
	switch a {
	case "a":
		if b == "b" {
			fmt.Println("555")
			break
		}
		fmt.Println("666")
	case "b":
		if aa == "aa" {
			fmt.Println("9999")
		}
		//fmt.Println("777")
	default:
		fmt.Println("888")
	}
}
