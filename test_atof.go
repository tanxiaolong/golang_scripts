package main

import "fmt"

func main() {
	str := "123.456"
	rst := 0.0    //存最终结果
	flag := false //标识是否有小数点
	temp := 0.1   //表示小数位数
	for _, v := range str {
		if v == '.' {
			flag = true
			continue
		}
		if !flag {
			rst = rst*10.0 + float64(v-'0')
		} else {
			rst = rst + float64(v-'0')*temp
			temp = temp * 0.1
		}
	}
	fmt.Println(rst)
	fmt.Println(atof(str))
}

func atof(a string) float64 {
	t := 1
	flag := false
	rlt := 0
	for i := range a {
		if a[i] == '.' {
			flag = true
			continue
		}
		if flag {
			t = t * 10
		}
		rlt = rlt*10 + int(a[i]-'0')

	}
	return float64(rlt) / float64(t)
}
