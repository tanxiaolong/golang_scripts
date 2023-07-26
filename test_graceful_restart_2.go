package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	server   *http.Server
	listener net.Listener
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.Write([]byte("hello world from pid:" + strconv.Itoa(os.Getpid())))
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	f, err := tl.File()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("listener is not tcp listencer")
	}
	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	for {
		sig := <-ch
		log.Printf("signal: %v", sig)

		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			log.Printf("graceful shutdown")
		case syscall.SIGUSR2:
			log.Printf("reload")
			err := reload()
			if err != nil {
				log.Fatal("graceful restart error: %v", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload")
		}
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/hello", handler)
	server = &http.Server{Addr: ":9999"}
	var err error
	if *graceful {
		log.Print("main: Listening to existing file descriptor 3.")

		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	}
	if err != nil {
		log.Fatal("listener error : %v", err)
	}

	go func() {
		err = server.Serve(listener)
		log.Printf("server.Serve err: %v\n", err)
	}()

	signalHandler()
	log.Printf("signal end")
}
