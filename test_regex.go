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

	// email
	text = "123@qq.com"
	text = ""
	reg = regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)
	valid := reg.MatchString(text)
	fmt.Println(valid)

	uid := "302005"
	reg = regexp.MustCompile(`^[1-9][0-9]+(,[1-9][0-9]+)*$`)
	find := reg.MatchString(uid)
	fmt.Println("333333", find)

	str := "TG1010;TG102;"
	str = "TG101"
	str = "TG1011;"
	str = "TG101;TG10102;"
	str = "TG101;TG102。。;"
	reg = regexp.MustCompile(`^(TG\d+;)+$`)
	find = reg.MatchString(str)
	fmt.Println(find)

	isMatch, _ := regexp.MatchString(`^(TG\d+;)+$`, str)
	fmt.Println(isMatch)
}
