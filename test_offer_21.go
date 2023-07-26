package main

import "fmt"

func main(){
	fmt.Println(exchange([]int{2,4,6}))
}
func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		for left < len(nums)&& nums[left]&1 != 0{
			left ++
		}
		for right >= 0 && nums[right] &1 == 0  {
			right --
		}

		if left>=right {
			break
		}
		nums[left],nums[right]=nums[right],nums[left]
		left ++
		right --
	}
	return nums
}

func spiralOrder(matrix [][]int) []int {
	column := len(matrix)
	row := len(matrix[0])
	l := 0
	r := row-1
	u := 0
	d := column-1
	order := []int{}
	for l <= r && u <= d {
		for column = l; column <= r; column++ {
			order = append(order,matrix[u][column])
		}
		for row = u + 1; row <= d; row++ {
			order = append(order,matrix[row][r])
		}
		if l < r && u < d {
			for column = r - 1; column > l; column-- {
				order = append(order,matrix[d][column])
			}
			for row = d; row > u; row-- {
				order = append(order,matrix[row][l])
			}
		}
		l++
		r--
		u++
		d--

	}
	return order
}
