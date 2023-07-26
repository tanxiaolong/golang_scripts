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
	str := "country: \344\270\255\345\233\275 countryCode:156 province:\345\271\277\344\270\234\347\234\201 city:\345\271\277\345\267\236\345\270\202 districtCode:440100 frontIsp:\350\205\276\350\256\257\347\275\221\347\273\234 backboneIsp:\350\205\276\350\256\257\347\275\221\347\273\234"
	str = "\345\210\253\346\210\252\345\261\217\344\272\206\345\245\275\345\220\227\357\274\237"
	res := convertOctonaryUtf8(str)
	res2, _ := url.QueryUnescape(str)
	fmt.Println(res2)
	fmt.Println(res)
}
