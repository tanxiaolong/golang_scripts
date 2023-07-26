package main

import "fmt"
import "time"

func main() {
	m := make(map[int]int)
	go func() {
		for {
			//_ = m[1]
			for k, v := range m {
				fmt.Println(k, v)
			}
			fmt.Printf("%p\n", m)
		}
	}()
	go func() {
		for {
			//m[2] = 1
			n := map[int]int{1: 1}
			for i := 0; i < 100; i++ {
				n[i] = i
			}
			m = n
			time.Sleep(1)
		}
	}()

	select {}
}
