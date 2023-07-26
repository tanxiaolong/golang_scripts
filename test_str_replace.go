package main

import "fmt"
import "github.com/syyongx/php2go"

func main() {
	fmt.Println("vim-go")
	str := "hell wolrd"

	tt := php2go.StrReplace("e", "o", str, -1)
	fmt.Println(tt)
}
