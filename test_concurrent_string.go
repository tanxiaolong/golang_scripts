package main

import "fmt"
import "time"
import "strings"

const (
	first  = "what"
	second = ""
)

func main() {
	var s string

	go func() {
		flag := false
		for {
			if flag {
				s = first
			} else {
				s = second
			}
			time.Sleep(10)
			flag = !flag
		}
	}()

	for {
		ss := strings.Split(s, ",")
		fmt.Println(ss)
		time.Sleep(10)
	}
}
