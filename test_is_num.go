package main

import "fmt"
import "regexp"

func main() {
	pattern := "\\d+"
	str := ".124534"
	result, _ := regexp.MatchString(pattern, str)
	fmt.Println(result)
}
