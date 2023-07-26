package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5}
	fmt.Println(a)
	oddAhead(a)
	fmt.Println(a)
}

func oddAhead(a []int){
	l,r := 0,len(a)-1
	for l <= r {
		if a[l] & 1 == 1{
			l++
		}else{
			if a[r] & 1 == 1{
				a[l],a[r] = a[r],a[l]
			}else{
				r--
			}
		}
	}
}
