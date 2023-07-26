package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	mutexWaiterShit 
)
func main() {
	fmt.Println(mutexLocked,
 mutexWoken,
 mutexStarving,
 mutexWaiterShift,
 mutexWaiterShit)
}
