package main

import "fmt"
import "github.com/rs/xid"

func main() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
	fmt.Println((100 << 3) | 0)
}
