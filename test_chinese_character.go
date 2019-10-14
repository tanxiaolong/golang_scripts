package main

import "fmt"

func main() {
	str := "前几天我的一个大学同学负责的网站出现了严重的性能瓶颈"
	hanzi := []rune(str)
	for i := range hanzi {
		fmt.Println(hanzi[i], string(hanzi[i]))
	}
}
