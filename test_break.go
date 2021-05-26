package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				break
			}

			fmt.Println(i)
		}
	}
}
