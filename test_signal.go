// go 1.10

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGUSR2,
		syscall.SIGTERM,
	)

	fmt.Println("pid is: ", os.Getpid())

	for {
		sig := <-signalChan
		fmt.Println("get signal: ", sig)
	}
}
