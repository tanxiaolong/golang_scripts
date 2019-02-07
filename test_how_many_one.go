package main

import "fmt"

func main(){
	fmt.Println(ones(9))	
}

func ones(target int) int{
	count := 0
	for target!=0 {
		count++
		target = (target-1)&target
	}
	return count
}
