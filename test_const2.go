package main

import "fmt"

const (
	i = 7
	j
	k
)

func main(){
	fmt.Println(i,j,k)
	a:=[]int{1,2,3,4,5,6,7}
	s := a[2:4]
	n := append(s, 55,66)
	fmt.Println(len(n),cap(n))
}
