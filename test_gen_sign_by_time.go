package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	tNow := time.Now().UnixNano() / 1e6
	fmt.Println(tNow)
	crypSign := GenSign(tNow)
	fmt.Println(crypSign)
}
func GenSign(reqTime int64) string {
	key := "OFFLINE_PUSH_" + fmt.Sprintf("%d", reqTime)
	h := md5.New()
	h.Write([]byte(key)) // 需要加密的字符串
	return hex.EncodeToString(h.Sum(nil))[0:16]
}
