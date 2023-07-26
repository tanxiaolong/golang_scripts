package main

import "fmt"
import "math"

func main() {
	rlt := math.Pow(2, 53)
	fmt.Println(fmt.Sprintf("%.0f", rlt-1))
}
