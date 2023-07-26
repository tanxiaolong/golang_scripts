package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func play(table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Println(ball, getGID())
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func getGID() uint64 {
	b := make([]byte, 64)
	fmt.Println(string(runtime.Stack(b, false)))
	b = b[:runtime.Stack(b, false)]
	fmt.Println(string(b))
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func main() {
	table := make(chan int)
	ball := 0
	go play(table)
	go play(table)
	table <- ball
	time.Sleep(1 * time.Second)
	<-table
}
