package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func checkSign(secret, timestamp, nonce, sign string) bool {
	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	// 如果请求头时间与当前时间相差超过180秒，则认为请求不合法 (结合x-rio-nonce，可以防止请求重放)
	now := time.Now().Unix()
	if ts > now+180 || ts < now-180 {
		log.Printf("time expired")
		return false
	}
	signData := fmt.Sprintf("%s%s%s%s", timestamp, secret, nonce, timestamp)
	log.Printf("signdata: %s", signData)
	res := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signData))))
	log.Printf("signres: %s", res)
	return res == sign
}

func main() {
	secret := "Nu4WmXGBauU99NsNuwXjbV7hO2xwEef91Rp9rKX6WatcmiwWqXyk"
	secret = "KQJhACtOaqm9qxgtmocpR4kRd0MJitHJ"
	headerTimestamp := "1523340394"
	headerTimestamp = fmt.Sprintf("%d", time.Now().Unix())
	headerSign := "B6CD9D9377EDB58CA445B4A6A753F884C156DBD6344F058E7B88658E9AF54F71"
	headerNonce := "ab570e0a:015a1d0fbef0:00cb3d"
	if checkSign(secret, headerTimestamp, headerNonce, headerSign) {
		log.Printf("signsuccesss")
	}
}
