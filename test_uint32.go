package main

import "fmt"
import "math"

const aa = 1610693105
const cc uint64 = 1610693105244

func main() {
	//fmt.Println(math.MaxUint64)
	//fmt.Println("max uint64", math.MaxUint64)
	fmt.Println("max uint32", math.MaxUint32)
	fmt.Println("max int32", math.MaxInt32)
	// case 1
	u32 := uint32(123)
	fmt.Println(u32)
	// case 2
	fmt.Println(uint32(aa))
	// case 3
	aaa := 1610693105
	fmt.Println(uint32(aaa))
	//a := 1610693105244
	//fmt.Println(uint32(a))
	//fmt.Println(a & 0xFFFFFFFF)
	//u32 = uint32(a)
	//fmt.Println(u32)
	//  constant 1610693105244 overflows uint32
	//fmt.Println(uint32(aa))
	//fmt.Println(uint32(cc))
	//fmt.Println(uint32(1610693105244))
	//fmt.Println(uint32(math.MaxUint64))

	var a uint32
	var b int32 = 3
	a = 2 - uint32(b)
	fmt.Println(a)
}
