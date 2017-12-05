package main

import "fmt"

func t(b ...interface{}){
	for i,j :=range b{
		fmt.Println(i,j)
	}
}
func main(){
	t([]interface{}{1,2,3,"4"}...)	
}
