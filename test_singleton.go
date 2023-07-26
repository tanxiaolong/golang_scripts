package main

import "fmt"

func main() {
	fmt.Printf("%p\n", s)
	ss := NewSingle()
	fmt.Printf("%p\n", ss)

	sss := NewSingle()
	fmt.Printf("%p\n", sss)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			ss := NewSingle()
			fmt.Printf("%p\n", ss)
		}(i)
	}
}

var s *single

type single struct {
	a int
}

func NewSingle() *single {
	if s != nil {
		return s
	}
	s = &single{a: 1}
	return s
}
