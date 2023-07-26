package main

import "fmt"
import "time"

func main() {
	a := []int{1, 23, 4, 5, 6}
	fmt.Println(a[0:3])
	fmt.Println(a[3:])
	var b int64 = 101234
	fmt.Println(b)

	now := time.Now()
	fmt.Println("unix:", now.Unix())
	fmt.Println("nano:", now.UnixNano())

}
