package main

import "unsafe"
import "fmt"

type User struct {
 A int32 // 4
 B []int32 // 24
 C string // 16
 D bool // 1
}

func main()  {
 var u User
 //u.B = []int32{1}
 //u.A = 1
 //u.C = "12"
 //u.D = true
 fmt.Println("u1 size is ",unsafe.Sizeof(u))
 fmt.Println(unsafe.Alignof(u))
 fmt.Println("aaa")
 fmt.Println(&u)
 fmt.Println(&u.A, unsafe.Alignof(u.A), unsafe.Sizeof(u.A))
 fmt.Printf("%p %d %d\n",&u.B, unsafe.Alignof(u.B), unsafe.Sizeof(u.A))
 fmt.Println(&u.C, unsafe.Alignof(u.C), unsafe.Sizeof(u.A))
 fmt.Println(&u.D, unsafe.Alignof(u.D), unsafe.Sizeof(u.A))
}


