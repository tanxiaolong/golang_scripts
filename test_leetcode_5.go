package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

*/
func main() {
	fmt.Println(longestPalindromeDP("aaaa"))
}

// dp
func longestPalindromeDP(s string) string {

	length := len(s)
	if length <= 1 {
		return s
	}
	var rlt string

	dp := map[int]map[int]bool{}

	for r := range s {
		//for r := 1; r < length; r++ {
		fmt.Println("r:", r)
		for l := 0; l < r; l++ {
			fmt.Println("l:", l)
			fmt.Printf("before:%+v\n", dp)
			if s[l] == s[r] && (dp[l+1][r-1] || r-l <= 2) {
				dp[l] = map[int]bool{r: true}
				fmt.Println("before rlt:", rlt)
				if len(rlt) < len(s[l:r+1]) {
					rlt = s[l : r+1]
					fmt.Println("after rlt:", rlt)
				}
			}
			fmt.Printf("after:%+v\n", dp)
		}
	}
	if rlt == "" {
		rlt = string(s[0])
	}
	return rlt
}

// 先马拉车，然后从中间向两边扩散比较
func longestPalindrome(s string) string {

	var rlt string

	var newS string
	var newSByteArr []byte
	newSByteArr = append(newSByteArr, '#')
	for i := range s {
		newSByteArr = append(newSByteArr, s[i])
		newSByteArr = append(newSByteArr, '#')
	}
	newS = string(newSByteArr)
	fmt.Println("str:", "0123456789")
	fmt.Println("str:", newS)
	length := len(newS)
	for i := 1; i < length-1; i++ {
		for m, n := i-1, i+1; m >= 0 && n < length; {
			fmt.Println("m:", m, "n:", n)
			if newS[m] == newS[n] {
				if len(rlt) < len(newS[m:n+1]) {
					rlt = newS[m : n+1]
				}
				m--
				n++
				continue
			}
			break
		}
		fmt.Println("break", rlt)
	}
	var bytesArr []byte
	for i := range rlt {
		if rlt[i] != '#' {
			bytesArr = append(bytesArr, rlt[i])
		}
	}
	rlt = string(bytesArr)
	return rlt
}
