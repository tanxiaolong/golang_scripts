package main

import "fmt"
import "time"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lnums1 := len(nums1)
	lnums2 := len(nums2)
	length := lnums1 + lnums2

	var curPos int
	var i, j int
	if length%2 == 0 {
		stopPos1, stopPos2 := length/2-1, length/2
		var val1, val2 float64
		val1, val2 = -1, -1
		for i < lnums1 || j < lnums2 {
			time.Sleep(time.Second)
			fmt.Println(nums1[i], nums2[j])
			if nums1[i] < nums2[j] {
				if curPos == stopPos1 {
					val1 = float64(nums1[i])
					break
				}
				if curPos == stopPos2 {
					val2 = float64(nums2[i])
					break
				}
				curPos++
				i++
			} else {
				if curPos == stopPos1 {
					val1 = float64(nums2[j])
					break
				}
				if curPos == stopPos2 {
					val2 = float64(nums2[j])
					break
				}
				curPos++
				j++
			}
		}

		fmt.Println("las:", val1, val2)
		fmt.Println("i:", i, "j:", j)
		return (val1 + val2) / 2
	} else {
		stopPos := length / 2
		for i < lnums1 && j < lnums2 {
			if nums1[i] < nums2[j] {
				if curPos == stopPos {
					return float64(nums1[i])
					break
				}
				curPos++
				i++
			} else {
				if curPos == stopPos {
					return float64(nums2[j])
					break
				}
				curPos++
				j++
			}
		}
	}
	return 0
}
