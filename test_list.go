package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(3)
	fmt.Println(l.Front().Value)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
