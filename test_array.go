package main

import "fmt"

func main(){
b := make([]int, 1,2 )	
a := make([]int, 10)
fmt.Println(len(b),cap(b), len(a),cap(a))
	
}
