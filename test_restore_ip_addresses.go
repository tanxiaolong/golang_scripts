package main


import "fmt"

type Ip_segment struct {
	first string
	second string
	third string
}

type Ip struct {
	ip_fmt [4]ip_segment
}

func restoreIp(ip string) []string{
	rlt := []Ip{}
	for _,j:= range ip{
		if rlt.ip_fmt.first == nil && string(j)<="2" && string(j) >0{
			rlt.ip_fmt.first = string(j)
		}else if rlt.ip_fmt.second == nil && string(j) <= "5" && string(j) > 0{
			rlt.ip_fmt.second = string(j)
		}else if rlt.ip_fmt.third == nil && string(j) <= "5" && string(j) > 0 {
			rlt.ip_fmt.third = string(j)
		}
		if rlt.ip_fmt.first
	}
	return rlt
}

func main(){
	ip := "25525511135"
	fmt.Println(restoreIp(ip))
}
