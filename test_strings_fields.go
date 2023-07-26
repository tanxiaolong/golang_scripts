package main

import "fmt"
import "strings"

func main() {
	fields := strings.Fields("0 12 * * *")

	fmt.Println(fields)
}
