package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "123,01239810"

	//查找连续的小写字
	reg := regexp.MustCompile(`^[1-9][0-9]+(,[1-9][0-9]+)+$`)
	foundBool := reg.MatchString(text)
	fmt.Println(foundBool)
}
