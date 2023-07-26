package main

import "fmt"
import "reflect"

func main() {
	a := map[int][3]int{}
	fmt.Println(a)
	a[0] = [3]int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a[0][0])
	//a[0][0] = 4
	fmt.Println(a)
	b := [2]int{1, 2}
	fmt.Println(reflect.TypeOf(b))
	c := []int{1, 2}
	fmt.Println(reflect.TypeOf(c))
	d := [...]int{2: 1, 4: 3}
	fmt.Println(d)
}
