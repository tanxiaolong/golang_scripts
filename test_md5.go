package main

import (
	"crypto/md5"
	//"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("Mdroid.cn")
	//md5Ctx := md5.New()
	//md5Ctx.Write(data)
	//cipherStr := md5Ctx.Sum(nil)
	//fmt.Println(cipherStr)
	fmt.Println(string(data))
	fmt.Printf("%x\n", md5.Sum(data))
	//fmt.Printf("%x\n", cipherStr)
	//fmt.Println(hex.EncodeToString(cipherStr))
}
