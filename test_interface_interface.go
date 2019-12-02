package main

import (
	"fmt"
)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	b := make([]byte,num)
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

type A struct {
	pos  int
	data []byte
}

func (a *A) next() byte {
	return a.data[a.pos]
}

/*
func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next() // 编译错误:n.next未定义（*nexter类型没有next成员或next方法）
	}
	return b
}
*/
func main() {
	a := &A{
		pos:  3,
		data: ([]byte)("hello world, welcome to China, my home country"),
	}
	fmt.Println(a.next())
	fmt.Println(nextFew1(a, 10))
}
