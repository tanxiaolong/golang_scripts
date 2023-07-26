package main

import "fmt"

type aa struct {
	a int
}

func main() {
	arr := []*aa{
		{a: 1},
		{a: 2},
		{a: 3},
		{a: 4},
		{a: 5},
	}
	afterDivide := divide(arr, 2)
	for k, v := range afterDivide {
		for i := range v {
			fmt.Printf("%d, %+v\n", k, v[i])
		}
	}
}

func divide(ori []*aa, size int) map[int][]*aa {
	if size <= 0 {
		size = 1
	}
	res := make(map[int][]*aa)
	idx := 0
	for i := 0; i < len(ori); i = i + size {
		if i+size < len(ori) {
			res[idx] = ori[i : i+size]
		} else {
			res[idx] = ori[i:]
		}
		idx++
	}
	return res
}
