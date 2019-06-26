package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := map[string]interface{}{
		"a": "b",
	}

	b := a
	a["b"] = "a"
	fmt.Println(a, b)

	c := map[string]interface{}{}
	for k, v := range a {
		c[k] = v
	}
	a["c"] = "c"

	fmt.Println(a, c)
}
