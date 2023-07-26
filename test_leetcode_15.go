package main

import "log"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums = []int{-1,0,1,2}
	nums = []int{13,-14,-10,-4,4,4,0,-14,5,-9,-3,-10,14,7,-3,-4,-3,12,-14,2,-11,-6,0,-7,13,-2,-7,-11,-14,-13,5,14,-12,11,-13,-1,-8,2,0,4,1,4,10,-8,-11,-8,3,1,11,4,-12,8,5,-4,-14,-4,9,-13,-8,2,-11,12,-7,14,0,-5,-2,7,5,5,-3,13,-6,-8,-10,-10,-9,0,6,-12,11,2,11,1,13,4,12,-1,6,-11,-14,2,9,-6,1,-6,-4,14,-13,8,4,-1,6,6,-2,11,11,4,-4,-5,-1,-8,12,-13,1,10,7,-10,-14,-10,-5,-13,0,11}
	log.Print(threeSum2(nums))
}
func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i,v := range nums {
        if p,exists := m[v];exists && target-v == v{
            return []int{nums[p],nums[i]}
        }
        m[v] = i
    }
    for i,v := range nums {
        left := target - v
        if left != v {
            if p,exists := m[left];exists {
                return []int{nums[i],nums[p]}
            }
        }
    }
    return []int{}
}

func threeSum2(nums []int) [][]int {
    rlt := [][]int{}
    tmp := map[int]int{}
    for i:=range nums {
	sad := []int{}
	sad = append(sad,nums[:i]...)
	sad = append(sad,nums[i+1:]...)
        tt := twoSum(sad,-nums[i])
        if len(tt) != 0 {
            t := append(tt,nums[i])
            sum := 0
            for _,v := range t {
                sum = sum+v*v
            }
            if _,exists := tmp[sum];!exists {
                rlt = append(rlt,t)
            }
	    tmp[sum] = 1
        }
    }
    return rlt
}

func qsort(a []int){
	if len(a) < 2 {
		return
	}
	l,r := 0,len(a)-1
	for i:=range a {
		if a[i] < a[r] {
			a[i],a[l] = a[l],a[i]
			l++
		}
	}
	a[l],a[r] = a[r],a[l]
	qsort(a[:l])
	qsort(a[l+1:])
}
func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return [][]int{}
	}
	rlt := [][]int{}
	smaller := []int{}
	bigger := []int{}
	for i := range nums {
		if nums[i] < 0 {
			smaller = append(smaller, nums[i])
		} else {
			bigger = append(bigger, nums[i])
		}
	}
	qsort(smaller)
	qsort(bigger)
	log.Print(smaller)
	log.Print(bigger)
	panic(1)
	ls := len(smaller)
	lb := len(bigger)
	if ls == 1 {
		for i := 0; i < lb; i++{
			for j := i + 1; j < lb; j++{
				if smaller[0]+bigger[i]+bigger[j] == 0 {
					rlt = append(rlt, []int{smaller[0], bigger[i], bigger[j]})
				}
			}
		}
		return rlt
	} else if lb == 1 {
		for i := 0; i < ls; i++{
			for j := i + 1; j < ls; j++{
				if bigger[0]+smaller[i]+smaller[j] == 0 {
					rlt = append(rlt, []int{bigger[0], smaller[i], smaller[j]})
				}
			}
		}
		return rlt
	}
	tmp := map[int]int{}
	for i := 0; i < ls; i++ {
		if _, exists := tmp[smaller[i]]; exists {
			continue
		} else {
			tmp[smaller[i]] = 1
		}
		j := i + 1
		for ; j < ls; j++ {
			for k := 0; k < lb; k++ {
				if smaller[i]+smaller[j]+bigger[k] == 0 {
					rlt = append(rlt, []int{smaller[i], smaller[j], bigger[k]})
				}
			}
		}
		//if j == ls-1 {
		for k := 0; k < lb; k++ {
			for h := k + 1; h < lb; h++ {
				if smaller[i]+bigger[k]+bigger[h] == 0 {
					rlt = append(rlt, []int{smaller[i], bigger[k], bigger[h]})
				}
			}
		}
		//}
	}

	return rlt
}
