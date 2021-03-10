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
	fmt.Println(5555, arr, len(arr))
	a := map[string]int{
		"a": 1,
	}
	for _, val := range arr {
		if _, ok := a[val]; !ok {
			fmt.Printf("1%s2\n", val)
		}
	}
	fmt.Println(123)

	ass := []string{"a", "b", "c", "d"}
	mark := 0
	for i, aa := range ass {
		if aa == "d" {
			mark = i
			break
		}
	}
	preHalf := ass[0:mark]
	fmt.Println(preHalf)
	newAss := append(preHalf, ass[mark+1:]...)
	fmt.Println(newAss)

	str = "   	, 	  ,c		"
	fmt.Println(str)
	str = strings.TrimSpace(str)
	fmt.Println(str)
	arr = strings.Split(str, ",")
	fmt.Printf("%q\n", arr)

	str = "     ,        "
	fmt.Println(str)
	str = strings.TrimSpace(str)
	fmt.Println(str)
	arr = strings.Split(str, ",")
	fmt.Printf("%q\n", arr)

	str = "     ,"
	arr = strings.Split(str, ",")
	fmt.Println("666", "???", arr[0] == "", "???", arr[1], "???", len(arr))
	str = ",     "
	arr = strings.Split(str, ",")
	fmt.Println("666", "???", arr[0] == "", "???", arr[1], "???", len(arr))
}
