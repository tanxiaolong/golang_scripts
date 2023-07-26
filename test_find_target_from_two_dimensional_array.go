package main

import "fmt"

func main() {
        arr := [][]int{
                {1, 2, 3, 4, 5, 6},
                {7, 8, 9, 10, 11, 12},
                {14, 15, 16, 17, 18, 19},
        }
        fmt.Println(search(arr, 13))
        /* 二分查找
        isExists := true
        for i := range arr {
                isExists = isExists || bsearch(arr[i],14)
        }
        fmt.Println(isExists)
        */
}
func bsearch(a []int, target int) bool{
        l,r := 0, len(a)-1
        for l <= r {
                mid := (l+r)/2
                if a[mid] == target {
                        return true
                }
                if a[mid] < target {
                        l = mid+1
                }
                if a[mid] > target {
                        r = mid-1
                }
        }
        return false
}
func search(a [][]int, target int) bool {
        col := len(a[0])
        row := len(a)
        j := 0
        for i := 0; i < row; {
                for j < col {
                        if j < 0 {
                                return false
                        }
                        if a[i][j] == target {
                                return true
                        }

                        if a[i][j] < target {
                                if (j + 1) < col {
                                        j++
                                } else if (i + 1) < row {
                                        i++
                                }
                                continue
                        }
                        if a[i][j] > target {
                                j--
                                continue
                        }

                }
        }
        return false
}

