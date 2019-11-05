package main

import (
	"log"
	"net"

	"github.com/panjf2000/gnet"
)




func (es *echoServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	c.ResetBuffer()
	return
}

func main() {
	echo := &echoServer{}

	echo.NumLoops = 1
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}