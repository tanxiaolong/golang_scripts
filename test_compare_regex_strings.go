package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	text := "0 12 * * *"
	count := 20000000

	tStart := time.Now().Unix()
	for i := 0; i < count; i++ {
		regexp.MustCompile(" ").Split(text, 5)
	}
	tEnd := time.Now().Unix()
	fmt.Println("regex: ", tEnd-tStart)

	tStart = time.Now().Unix()
	for i := 0; i < count; i++ {
		strings.Split(text, " ")
	}
	tEnd = time.Now().Unix()
	fmt.Println("string.Split: ", tEnd-tStart)

	tStart = time.Now().Unix()
	for i := 0; i < count; i++ {
		strings.Fields(text)
	}
	tEnd = time.Now().Unix()
	fmt.Println("strings.fields: ", tEnd-tStart)

}
