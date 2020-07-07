package main

import "fmt"
import "encoding/hex"

func main() {
	buf := make([]byte, 16)
	uuid := hex.EncodeToString(buf)
	fmt.Println(uuid)
}
