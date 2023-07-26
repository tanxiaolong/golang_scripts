package main

import "fmt"
import "time"

func main() {

	a := []int{3, 4, 5, 6, 7, 8}
	for i := range a {
		go func(i int) {
			fmt.Println(a[i])
		}(i)
	}

	time.Sleep(10)

	fmt.Println("================")
	for _, v := range a {
		go func(i int) {
			fmt.Println(i)
		}(v)
	}
	time.Sleep(10)

	fmt.Println("================")
	for i := 0; i < len(a); i++ {
		go func(i int) {
			fmt.Println(a[i])
		}(i)
	}

	time.Sleep(10)

}
