package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

func main() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()

	rec1, err := c1.Do("set", "a", 1)
	if rec1 != nil {
		fmt.Printf("---------  ")
		fmt.Println(rec1)
	}
	rec2, err := c1.Do("get", "a")
	if rec2 != nil {
		fmt.Printf("---------  ")
		fmt.Println(rec2)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			_, err = c1.Do("incr", "aaa", 1)
			if err != nil {
				panic(err.Error())
			}
		}(i)
	}
	wg.Wait()

	rec2, _ = c1.Do("get", "aaa")
	fmt.Printf("incr aaa rlt: %v\n", rec2)

}
