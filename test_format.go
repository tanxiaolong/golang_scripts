package main

import "fmt"

func main() {
	fmt.Printf("%02s\n", "5")
	fmt.Printf("%02s\n", "16")
	fmt.Printf("t_flow_coupon_record_%04d%02d\n", 2020, 9)

	fmt.Printf("剩余%.1f天%.1f小时\n", 0.55, 10)
	tmp := 10
	fmt.Printf("剩余%.1f天%.1f小时", 0.55, float32(tmp))

}
