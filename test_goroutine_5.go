
package main

import "time"
import "fmt"

func main() {
defer func() {
	err := recover()
fmt.Println(err)
}()
    ch := make(chan struct{})
    go func() {
        fmt.Println("start working")
        time.Sleep(time.Second * 1000)
        ch <- struct{}{}
    }()

    <-ch

    fmt.Println("finished")
}
