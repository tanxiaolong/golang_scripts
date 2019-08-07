package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func main() {
	var users []map[string]interface{}
	var user = make(map[string]interface{})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var names = []string{
		`calvin`,
		`jason`,
		`bob`,
		`hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu`,
	}
	fmt.Println("size of name:", names[0], unsafe.Sizeof(names[0]))
	fmt.Println("size of name:", names[1], unsafe.Sizeof(names[1]))
	fmt.Println("size of name:", names[2], unsafe.Sizeof(names[2]))
	fmt.Println("size of name:", names[3], unsafe.Sizeof(names[3]))

	//for _, name := range names {
	for i := 0; i < 1000000; i++ {
		name := names[r.Intn(2)]
		user[`name`] = name

		users = append(users, user)
		// always is 24
		//fmt.Println("after add name:", name, unsafe.Sizeof(users))
	}
	fmt.Println(cap(users))

}
