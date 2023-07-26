package main


import (
	"fmt"
)


func main(){
	var a = make([]int,1,2)
	fmt.Printf("%d,%d,%v\n",len(a),cap(a),a)
	//a[1] = 2
	//fmt.Println(a[1])
	a = append(a,1)
	fmt.Println(a)
	a = append(a,2)
	fmt.Println(a)
	var b = new([]int) // new 返回一个大小为0的类型指针
	//fmt.Printf("%d,%d\n",len(b),cap(b)) // 这个会报错，因为b是一个指针，没有len和cap方法
	fmt.Printf("%d,%d\n",len(*b),cap(*b))
	fmt.Println(*b)
	//(*b)[0] = 1 // 报错，因为没有初始化slice的大小
	*b = append(*b,1)
	fmt.Println(*b)
}
