package main

import (
	"fmt"
	"hash/crc64"
	"reflect"
)

func main() {
	s := "打死udhanckhdkja"
	s = "49d3f976734b089fe0ba28960948459d"
	//先建立一個table
	table := crc64.MakeTable(crc64.ECMA)
	//傳入位元組切片和table，返回一個uint64
	fmt.Println(crc64.Checksum([]byte(s), table)) //4295263180068867775
	s = "f0a8bbfdcf08874262abb3aeabdeaa69"
	rlt := crc64.Checksum([]byte(s), table)
	fmt.Println(reflect.TypeOf(rlt))

}
