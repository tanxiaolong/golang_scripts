package main

import (
	"fmt"
	"hash/crc32"
)

func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func main() {
	str := "123456abcd"
	str = "49d3f976734b089fe0ba28960948459d"
	hc := String(str)
	fmt.Println("hashcode:", hc)

	str = "f0a8bbfdcf08874262abb3aeabdeaa69"
	hc = String(str)
	fmt.Println("hc:", hc)

	// todo the two string has same crc32 result....
	// please refer test_crc64.go
}
