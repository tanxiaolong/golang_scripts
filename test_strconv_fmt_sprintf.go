package main

import "fmt"
import "strconv"
import "time"

func main() {
	t1 := time.Now()
	fmt.Println(strconv.Itoa(1))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	t3 := time.Now()
	fmt.Sprintf("%d", "1")
	t4 := time.Now()
	fmt.Println(t4.Sub(t3))
	a, err := strconv.Atoi("")
	fmt.Println(a, err)
}
