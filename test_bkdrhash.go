package main

import "fmt"

func main() {
	str := "b36a818dwslive1_%d664246ed"
	a := 1360271610
	sum := map[int32]int{0: 0, 1: 0, 2: 0, 3: 0}
	for i := 0; i < 10000; i++ {
		b := a + i
		str = fmt.Sprintf(str, b)
		mod := BKDRHash(str) % 4
		sum[mod]++
	}
	fmt.Println(sum)
}
func BKDRHash(str string) int32 {
	if len(str) == 0 {
		return 0
	}
	seed := int32(131)
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		b := str[i]
		hash = seed*hash + int32(b)
	}
	return hash & 0x7FFFFFFF
}
