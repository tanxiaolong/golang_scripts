package main

import "fmt"
import "math/rand"

func main() {
	k := 10
	hashes := make([]int, k)
	for i := 0; i < k; i++ {
		a := uint(rand.Int())
		b := uint(rand.Int())
		c := uint(rand.Int())
		x := computeHash(a*b*c, a, b, c)
		hashes[i] = int(x)
	}
	fmt.Println(hashes)
}

func computeHash(x, a, b, u uint) uint {
	return (a*x + b) >> (32 - u)
}
