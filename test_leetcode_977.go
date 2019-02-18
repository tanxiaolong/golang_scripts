package main

import (
	"log"
	"sort"
)

func main() {
	a := []int{-4, -1, 0, 3, 10}
	a = []int{-7, 2, 3, 11}
	a = []int{0, 2}
	a = []int{-3, 0, 2}
	//a = []int{0,0,2}
	//log.Print(sortedSquares(a))
	log.Print(so(a))
}

func so(A []int) []int{
	flag := 0
        for i := range A{
		A[i] = A[i]*A[i]
        }
	sort.Sort(A)
	return A
}
func sortedSquares(A []int) []int {
	if len(A) == 0 {
		return A
	}
	if len(A) == 1 {
		return []int{A[0] * A[0]}
	}
	if len(A) == 2 {
		a := A[0] * A[0]
		b := A[1] * A[1]
		if a > b {
			return []int{b, a}
		}
		return []int{a, b}
	}
	rlt := []int{}
	flag := 0
	for i, val := range A {
		A[i] = val * val
		if i-1 >= 0 && flag == 0 && A[i] >= A[i-1] {
			flag = i - 1
		}
	}
	log.Print(flag)
	bigger := A[:flag]
	smaller := A[flag:]
	log.Print(bigger,smaller)
	b := flag
	s := len(A) - flag
	i, j := 0, b-1
	for i < s && j >= 0 {
		if smaller[i] < bigger[j] {
			rlt = append(rlt, smaller[i])
			i++
			continue
		}
		rlt = append(rlt, bigger[j])
		j--
	}
	log.Print(i,j)
	if j >= 0 {
		rlt = append(rlt, bigger[j:]...)
	}
	rlt = append(rlt, smaller[i:]...)
	return rlt
}
