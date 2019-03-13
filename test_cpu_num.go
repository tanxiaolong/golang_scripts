package main

import (
    "runtime"
    "log"
)

func main(){
    log.Print(runtime.NumCPU())
}
