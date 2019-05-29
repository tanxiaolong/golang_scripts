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

	//str = "1,2,3"
	//arr = strings.Split(str, ",")
	str = " "
	arr = strings.Split(str, ",")
	fmt.Println(5555)
	a := map[string]int{
		"a": 1,
	}
	for _, val := range arr {
		if _, ok := a[val]; !ok {
			fmt.Printf("1%s2\n", val)
		}
	}
	fmt.Println(123)
}
