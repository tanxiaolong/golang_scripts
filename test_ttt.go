package main

import "fmt"

var Gmap map[int]int

func main() {

	Gmap = make(map[int]int)

	go func() {
		tmpMap := make(map[int]int)
		for i := 0; i < 100; i++ {
			tmpMap[1] = 1

		}

		Gmap = tmpMap
	}()

	go func() {
		for i := 0; i < 100; i++ {
			_ = Gmap[1]
		}
	}()

	fmt.Println(Gmap)

}
