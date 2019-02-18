package main

import "fmt"


func main(){
	a := []int{1,2,3,4,5}
	fmt.Println(a)
	switchOddAhead(a)
	fmt.Println(a)
}

func switchOddAhead(a []int){
	length := len(a)
	for i:=0; i< length/2; i++{
		f := 2*i + 1
		tmp := a[i]
		a[i] = a[f]
		a[f] = tmp
	}
}
