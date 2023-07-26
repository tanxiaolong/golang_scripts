package main

import (
    "fmt"
    "os/signal"
    "os"
    "syscall"
)

func handleSignal(signal os.Signal){
    fmt.Println("Got ",signal)
}

func main(){
    sigs := make(chan os.Signal, 1)

    signal.Notify(
        sigs,
        os.Interrupt,
        syscall.SIGHUP,
        syscall.SIGUSR2,
        syscall.SIGTERM,
    )
    fmt.Println(os.Getpid())
    go func(){
        for {
            sig := <- sigs
            fmt.Println(sig)
            if sig == os.Interrupt {
                os.Exit(-1)
            }
            handleSignal(sig)
        }
    }()
    a := make(chan int)
    <- a
}
