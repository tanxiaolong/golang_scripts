package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var g atomic.Value
	g.Store(map[int]interface{}{1: map[int]int{}})
	var b map[int]interface{}
	b, ok := g.Load().(map[int]interface{})
	if !ok || len(b[1].(map[int]int)) == 0 {
		fmt.Println(ok, len(b[1].(map[int]int)))
	}
	fmt.Println(b, ok)
	cc, ok := b[1].(map[int]int)
	fmt.Println(cc, ok)
}
