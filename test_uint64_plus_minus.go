package main

import "fmt"

func main() {
	var a uint64 = 3
	var b int64 = 10
	var c uint32 = 1
	fmt.Println("a:", a)
	fmt.Println("b:", uint64(b))
	fmt.Println("c:", uint64(c))
	fmt.Println("uint64 a-b:", a-uint64(b))
	fmt.Println("uint64 b-a:", uint64(b)-a)
	fmt.Println("int64 a-b:", int64(a)-b)
	if a-uint64(b) < uint64(c) {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

	if int64(a)-b < int64(c) {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}

	var d uint32 = 19
	var e uint32 = 18
	fmt.Println("uint32 d-e:", d-e)
	fmt.Println("uint32 e-d:", e-d)
	fmt.Println("d:", d)
	fmt.Println("uint32 c/d: ", c*100/d)
	TR := fmt.Sprintf("%.2f", float64(c)/float64(d))
	fmt.Println("c/d: ", TR, float64(c)/float64(d)*100)
}
