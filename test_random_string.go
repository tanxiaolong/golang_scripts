// go 1.10

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var bit int = 10
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	rlt := []byte{}
	for i := 0; i < bit; i++ {
		rlt = append(rlt, charset[r.Intn(len(charset))])
	}
	fmt.Println(string(rlt))

	ranNumber := 100000000 + r.Intn(999999999)
	fmt.Println(ranNumber)
}
