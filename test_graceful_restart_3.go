package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	server := &http.Server{Addr: "0.0.0.0:8888"}

	var gracefulChild bool
	var l net.Listener
	var err error

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")

	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)

		parent := syscall.Getpid()
		log.Printf("main: killing parent pid: %v", parent)
		syscall.Kill(parent, syscall.SIGTERM)
	} else {
		log.Println("main: Listening on a new file descriptor.")
		l, err = net.Listen("tcp", server.Addr)
	}
	server.Serve(l)
}
