package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := make([]int,10)
	a[1] = 1
	a[2] =3
	fmt.Println(len(a),cap(a))
}
