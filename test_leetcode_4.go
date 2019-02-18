package main

import "log"

func main(){
    log.Print(5*1.0/2)
    log.Print(findMedianSortedArrays([]int{1,2},[]int{3,4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := merge(nums1,nums2)
    l := len(arr)
    if l&1 == 0 {
	log.Print(float64(arr[l/2-1]+arr[l/2])/2)
        return float64(((arr[l/2-1]+arr[l/2])*1.0)/2)
    }else{
        return float64(arr[l/2])
    }
}
func merge(left, right []int) []int {
        result := make([]int, 0)
        m, n := 0, 0 // left和right的index位置
        l, r := len(left), len(right)
        for m < l && n < r {
                if left[m] > right[n] {
                        result = append(result, right[n])
                        n++
                        continue
                }
                result = append(result, left[m])
                m++
        }
        result = append(result, right[n:]...)
        result = append(result, left[m:]...)
        return result
}

