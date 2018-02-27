package main

import "fmt"

func foo() string {
	panic("异常")
	return "结束"
}

func main() {
	fmt.Println(foo())
}
