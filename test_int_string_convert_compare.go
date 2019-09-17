package main

import "fmt"
import "strconv"
import "time"

func main() {
	b := ""
	bInt, _ := strconv.Atoi(b)
	fmt.Println("bint:", bInt)
	a := 10000000
	start := time.Now().UnixNano()
	for i := 0; i < a; i++ {
		fmt.Sprintf("%d", a)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)

	start = time.Now().UnixNano()
	for i := 0; i < a; i++ {
		strconv.Itoa(a)
	}
	end = time.Now().UnixNano()
	fmt.Println(end - start)
}
