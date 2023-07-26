package main

import "fmt"
import "reflect"

func main() {
	a := 2450
	fmt.Println(reflect.TypeOf(a / 1000.0))
	fmt.Println(reflect.TypeOf(float64(a) / 1000))
	fmt.Println(reflect.TypeOf(2450 / 1000.0))
	fmt.Println(2450 / 1000.0)
	fmt.Println(2.4 / 10)
	fmt.Println(10 / 2.4)
	fmt.Println(fmt.Sprintf("%.2f", 10/2.4))
}
