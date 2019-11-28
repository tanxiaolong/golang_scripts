package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := map[interface{}]interface{}{
		"1": "123",
	}

	var b string = "1"
	if a[b] != nil {
		fmt.Println(a[b])
	}
}
