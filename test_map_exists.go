package main

import "fmt"

func main() {
		a:= map[int]int{
			1:1,
		}
	if val,exists := a[1];exists {
		fmt.Println(val)
	}
	fmt.Println(a[2])
}
