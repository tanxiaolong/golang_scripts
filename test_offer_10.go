package main

import "fmt"

func main(){
	fmt.Println(numWays(2))
}
func numWays(n int) int {

	return frog(n)
}
func frog(n int ) int {
	if n == 1 || n == 2{
		return 1
	}
	return frog(n-1)+frog(n-2)
}