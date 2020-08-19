package main

import "fmt"
import "time"
import "math"

func main() {

	fmt.Println(int64(math.Round(float64(time.Now().UnixNano()) / 1e6)))
}
