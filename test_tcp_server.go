package main

import (
	"fmt"
	"net"
)

func main() {
	lfd, err := net.Listen("tcp", ":8008")
	lfd.Accept()
	fmt.Println(err)
}
