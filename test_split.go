package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "1,2,3,4" //case 1
	arr := strings.Split(str, ",")
	fmt.Printf("%q\n", arr)
	fmt.Println(len(arr))
	str = "" //case 2
	arr = strings.Split(str, ",")
	fmt.Printf("%q\n", arr)
	fmt.Println(len(arr))
}
