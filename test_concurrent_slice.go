package main

import (
	"fmt"
)

func main(){
	a := []int{}
	for i:=0;i<1000;i++{
		a =append(a,i)
	}
	fmt.Println(len(a))
	b := make([]int,len(a))
	for i := range a {
		go func(i int) {
			if i%2 == 0 {
				return
			}
			b[i] = a[i]*a[i]
		}(i)
	}
	fmt.Println(b)
	fmt.Println(len(a) == len(b))
}
