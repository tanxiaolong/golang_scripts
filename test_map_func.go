package main

import "fmt"

type Msg func(name string) string

func main() {
	a := map[int]Msg{
		1: handle1,
	}

	fmt.Println(a[1]("tanxiaolong"))
}

func handle1(name string) string {
	fmt.Println(name)
	return "handle1"

}
