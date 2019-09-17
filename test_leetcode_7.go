package main

import "fmt"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

*/
func main() {
	fmt.Println("vim-go")
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	var f int //0+ 1-
	if x < 0 {
		f = 1
		x = -x
	}
	// 2147483647
	INT_MAX := int(^uint32(0) >> 1)

	//-2147483648
	INT_MIN := ^INT_MAX

	numbers := []int{}
	for x != 0 {
		a := x % 10
		numbers = append(numbers, a)
		x = x / 10
	}
	length := len(numbers)
	var rlt int
	if length != 0 {
		if length == 10 && numbers[0] > 2 {
			return rlt
		} else {
			for i := range numbers {
				var l int = 1
				for j := 1; j < length; j++ {
					l = l * 10
				}
				rlt = rlt + numbers[i]*l
				length--
			}
		}
	}
	if f == 1 {
		rlt = -rlt
		if rlt < INT_MIN {
			return 0
		}
	}
	if rlt > INT_MAX {
		return 0
	}
	return rlt
}
