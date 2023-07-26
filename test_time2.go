package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	extraHour := 9
	t2 := t.Add(3600 * time.Duration(extraHour) * time.Second)
	fmt.Printf("%v", t2)
}
