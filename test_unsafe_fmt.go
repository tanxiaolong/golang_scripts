package main

import (
	"fmt"
	"log"
	"signal"
)

func main() {
	log.Print("a")
	log.Print("b")
	syscall.Notify(15)
	fmt.Println("c")
	fmt.Println("d")
}
