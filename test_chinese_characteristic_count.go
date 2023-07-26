package main

import "fmt"
import "unicode"
import "unicode/utf8"
import "regexp"

func main() {
	sentence := "合利屋可以取餐了hahaha"
	fmt.Println(sentence)
	fmt.Println(len(sentence))
	l := utf8.RuneCountInString(sentence)
	fmt.Println("length of sentence is:", l)
	for _, word := range sentence {
		if unicode.Is(unicode.Scripts["Han"], word) {
			fmt.Println(word)
		} else {
			fmt.Println("not chinese.")
		}
	}

	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	fmt.Println(hzRegexp.MatchString(sentence))
}
