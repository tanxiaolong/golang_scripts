package main

import "fmt"
import "regexp"
import "strconv"

func main() {
	fmt.Println("vim-go")
	fmt.Println(converOctonaryUTF8("346200241"))
}

func converOctonaryUTF8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s, func(b []byte) []byte {
		i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
		return []byte{byte(i)}
	},
	)
	return string(out)
}
