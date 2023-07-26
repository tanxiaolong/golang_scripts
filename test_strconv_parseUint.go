//go 1.10
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseUint("1234", 10, 10))
	fmt.Println(strconv.ParseUint("1234", 10, 11))
}
