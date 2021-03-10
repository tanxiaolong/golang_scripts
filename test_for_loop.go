package main

import "fmt"
import "time"

func main() {
	n := 0
	for {
		time.Sleep(time.Second)
		n++
		fmt.Println(n)
		if n != 10 {
			continue
		}
		func() {
			fmt.Println("do sth")
		}()
	}
	nums := []int{1, 2, 3, 4, 5, 6}
	for i, num := range nums {
		if i+1 < 1 {
			panic("asd")
		}
		fmt.Println(i, num)
	}
}
