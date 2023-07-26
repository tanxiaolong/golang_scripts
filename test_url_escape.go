package main

import "fmt"
import "net/url"

func main() {
	str := "a\"b"
	fmt.Println(str)
	str = url.PathEscape(str)
	fmt.Println(str)
}
