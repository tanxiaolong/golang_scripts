package main

import (
	"fmt"
	"os"
	"sync"
)

/*
outputs:

9 10 0xc000084010 0xc000084018
1 2 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018
9 10 0xc000084010 0xc000084018

0 1 0xc000084060 0xc000084068
9 10 0xc0000c4018 0xc0000c4020
5 6 0xc000084078 0xc000084080
2 3 0xc000084098 0xc0000840a0
3 4 0xc0000840b8 0xc0000840c0
6 7 0xc0000c4038 0xc0000c4040
7 8 0xc0000c4058 0xc0000c4060
4 5 0xc0000840d8 0xc0000840e0
8 9 0xc0000c4078 0xc0000c4080
1 2 0xc0000a4018 0xc0000a4020



*/
func main() {

	collection := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg1 := sync.WaitGroup{}
	wg1.Add(10)
	for key, value := range collection {
		fmt.Println("out:", key, value, &key, &value)
		go func() {
			defer wg1.Done()
			fmt.Println("in:", key, value, &key, &value)
		}()
	}
	wg1.Wait()

	wg2 := sync.WaitGroup{}
	wg2.Add(10)
	fmt.Println()
	for key, value := range collection {
		fmt.Println("out:", key, value, &key, &value)
		go func(i, j int) {
			defer wg2.Done()
			fmt.Println("in:", i, j, &i, &j)
		}(key, value)
	}
	wg2.Wait()

}
