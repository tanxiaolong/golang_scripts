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

	pattern := `^[1-9][0-9]+$`
	str = "124534000000000000.0000"
	result, err := regexp.MatchString(pattern, str)
	fmt.Println("\\d+", result, err)

	// pics
	pattern = `(.*)\.(jpg|bmp|gif|ico|pcx|jpeg|tif|png|raw|tga)$`
	str = "https://xxx.jpg"
	result, err = regexp.MatchString(pattern, str)
	fmt.Println("pics:", result, err)

	pattern = "\\d+" //反斜杠要转义
	str = "1245340981209381209381290319023810923810923810283012830192839128301298319238102931239102381902381902831028319283102831092830128318290"
	result, _ = regexp.MatchString(pattern, str)
	fmt.Println(result)
}
