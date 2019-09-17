package main

import "fmt"
import "unicode"
import "unicode/utf8"

func main() {
	fmt.Println("vim-go")
	sentence := "合利屋可以取餐了hahaha"
	l := utf8.RuneCountInString(sentence)
	fmt.Println("length of sentence is:", l)
	for _, word := range sentence {
		if unicode.Is(unicode.Scripts["Han"], word) {
			fmt.Println(word)
		} else {
			fmt.Println("not chinese.")
		}
	}
}
