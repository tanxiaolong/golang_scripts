package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	url := "http://127.0.0.1:8080/hello"
	go func() {
		i := 0
		for {
			if i == 5 {
				break
			}
			time.Sleep(200 * time.Millisecond)
			req(url)
			i++
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println(". server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
	rand.Intn()

}

func req(url string) {
	c := &http.Client{
		Timeout: 3 * time.Second,
	}
	rsp, err := c.Get(url)
	log.Println(rsp, err)
}
