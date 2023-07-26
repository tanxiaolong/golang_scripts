package main

import "fmt"
import 	"unicode/utf8"

func main(){
	s := "今天"
	fmt.Println("length bytes",len(s))
	fmt.Println("length chars", utf8.RuneCountInString(s))
}
