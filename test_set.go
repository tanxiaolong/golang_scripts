package main

import (
	"fmt"
	"unsafe"
)

func main() {
	variant1 := make(map[string]bool)
	variant2 := make(map[string]struct{})
	variant3 := make(map[string]*struct{})
	for i := 0; i < 1<<16; i++ {
		if i%2 == 0 {
			continue
		}
		key := fmt.Sprintf("%v", i)
		variant1[key] = true
		variant2[key] = struct{}{}
		variant3[key] = &struct{}{}
	}
	size1 := unsafe.Sizeof(variant1)
	size2 := unsafe.Sizeof(variant2)
	size3 := unsafe.Sizeof(variant3)
	for k, v := range variant1 {
		size1 += unsafe.Sizeof(k)
		size1 += unsafe.Sizeof(v)
	}
	for k, v := range variant2 {
		size2 += unsafe.Sizeof(k)
		size2 += unsafe.Sizeof(v)
	}
	for k, v := range variant3 {
		size3 += unsafe.Sizeof(k)
		size3 += unsafe.Sizeof(v)
	}

	fmt.Printf("bool variant  : \t%v bytes\n", size1)
	fmt.Printf("struct variant: \t%v bytes\n", size2)
	fmt.Printf("point struct variant: \t%v bytes\n", size3)
	// bool variant  : 1114120 bytes
	// struct variant: 1048584 bytes
}
