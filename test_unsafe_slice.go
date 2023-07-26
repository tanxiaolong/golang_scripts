package main

import (
 "fmt"
 "reflect"
 "unsafe"
)

func main() {

 var s1 []int
 s2 := make([]int,0)
 s4 := make([]int,0)
 
 fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)),*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
 fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
 fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)

 s2 = []int{1}
 s4 = []int{1}
 fmt.Printf("s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))

 s := "123"
 sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
 
 b := []byte{}
 bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
 bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
 defer func(){
	e := recover()
	fmt.Println(e)
 }()
 b[0] = '4'
 fmt.Println(b)
}
