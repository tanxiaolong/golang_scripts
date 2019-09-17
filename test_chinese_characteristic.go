package main

import "fmt"
import "unicode"

func main() {
	fmt.Println("vim-go")
	sentence := "合利屋可以取餐了hahaha"
	for _, word := range sentence {
		if unicode.Is(unicode.Scripts["Han"], word) {
			fmt.Println(word)
		} else {
			fmt.Println("not chinese.")
		}
	}
}
