package main

import "fmt"
import "sync"

func main() {

	a := sync.Map{}
	vv, ok := a.Load("1")
	fmt.Println(vv, ok) //one true
	vv, ok = a.LoadOrStore("1", "one")
	fmt.Println(vv, ok) //one false
	vv, ok = a.Load("1")
	fmt.Println(vv, ok) //one true
}
