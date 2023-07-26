package main

import "fmt"

var GMap map[int]int

func main() {
	GMap = make(map[int]int)
	go func() {
		tmpMap := make(map[int]int)
		//for i := 0; i < 100; i++ {
		for {
			tmpMap[1] = 1
		}
		GMap = tmpMap
	}()

	go func() {
		//for i := 0; i < 100; i++ {
		for {
			_ = GMap[1]
			fmt.Printf("%p\n", GMap)
		}

	}()
	select {}
	fmt.Println(GMap)
}
