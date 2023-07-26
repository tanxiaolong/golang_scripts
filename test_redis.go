package main

import (
		"fmt"
		"github.com/garyburd/redigo/redis"
	"sync/atomic"
)

func main() {
	fmt.Println("vim-go")
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Printf("err: %v",err)
		return
	}

	var num int
	for i:=0;i<1000;i++ {
		go func(c redis.Conn){
			reply,err := c.Do("get","abc")
			fmt.Println(reply,err)
			//c.Do("set","abc",n)
		}(c)
	}

	fmt.Println(num)
}


