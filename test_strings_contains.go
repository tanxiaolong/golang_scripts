package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Contains("ZTT1.2.10_Android", "1.2.1"))
	fmt.Println(strings.Contains("ZTT1.2.10_Android", ""))
}
