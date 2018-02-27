package main

import "fmt"

// 如果1和2换了顺序，输出结果不一样
func main() {
	defer func() { // -----------------1
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { // ------------------2
		fmt.Println("1")
	}()
	panic("fault")
	fmt.Println("2")
}
