package main

import (
	"fmt"
	"time"
)

/*
遍历的时候，记录每个字符的位置，当有重复的时候回退到之前的i-(i-ilast-1)的位置，重新遍历
*/
func main() {
	fmt.Println(lengthOfLongestSubstring("c"))
}

func lengthOfLongestSubstring(s string) int {
	tmp := map[byte]int{}
	var length int = len(s)
	fmt.Println(length)
	var sum int
	var count int
	for i := 0; i < length; i++ {
		time.Sleep(time.Second)
		fmt.Println(string(s[i]), "current index:", i)
		if _, exists := tmp[s[i]]; !exists {
			tmp[s[i]] = i
			count++
		} else {
			if count > sum {
				sum = count
			}
			fmt.Println(string(s[i]), tmp[s[i]], i)
			i = i - (i - tmp[s[i]] - 1)
			fmt.Println("back index:", i)
			tmp = map[byte]int{}
			tmp[s[i]] = i
			count = 1
		}
	}
	if count > sum {
		sum = count
	}
	return sum
}
