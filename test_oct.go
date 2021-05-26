package main

import (
	"fmt"
	"net/url"
	"unicode/utf8"
)

func convertOctonaryUtf8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}
	return s
}

func main() {
	str := "\350\265\265\345\233\233"
	//res := convertOctonaryUtf8(http )
	res2, _ := url.QueryUnescape(str)
	fmt.Println(res2)
}
