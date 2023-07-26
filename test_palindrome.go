package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	str := "cabbaereweqwesds"
	res := ""
	for i := range str {
		s1 := palindrome(str, i, i)
		s2 := palindrome(str, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	fmt.Println(res)
}

func palindrome(str string, l, r int) string {
	for l >= 0 && r < len(str) && str[l] == str[r] {
		l--
		r++
	}

	return str[l+1 : r]
}
