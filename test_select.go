package main

import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 10
    }()

    go func() {
        ch2 <- 20
    }()

    select {
    case num := <-ch1:
        fmt.Println("Received from ch1:", num)
    case num := <-ch2:
        fmt.Println("Received from ch2:", num)
    }
}

