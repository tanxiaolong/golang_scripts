package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	fmt.Println(x, ",", x < 0)

	var arr []int
	for x != 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	length := len(arr)
	mid := length / 2
	var count int
	fmt.Println(arr)
	for i := 0; i < mid; i++ {
		fmt.Println(arr[i], ",", arr[length-1])
		if arr[i] == arr[length-1] {
			count++
		}
		length--
	}
	fmt.Println(count)
	if count != mid {
		return false
	}
	return true
}
