package main

import "fmt"

func kmp(origin, find string) bool {
	i := 0
	j := 0
	//lenOrigin := len(origin)
	lenFind := len(find)
	for i < len(origin) {
		if j < lenFind && find[j] == origin[i] {

			i++
			j++
		} else {
			i++
			break
		}
	}
	return false
}

func matchTable2(str string) map[int]int {
	mt := map[int]int{0: 0}

	for i, k := 1, 0; i < len(str); i++ {
		for k > 0 && str[i] != str[k] {
			k = mt[k-1]
		}

		if str[i] == str[k] {
			k++
		}
		mt[i] = k
	}
	fmt.Println(mt)
	return mt
}

func matchTable(str string) map[int]int {
	mt := map[int]int{0: 0}

	for i := 2; i < len(str)+1; i++ {
		sub := str[:i]
		// 获取所有前缀
		pres := findPresOrSufs(sub, true)
		// 获取所有后缀
		suffs := findPresOrSufs(sub, false)
		// 找到最长共有字符串
		max := 0
		for m, _ := range pres {
			if _, ok := suffs[m]; ok {
				if len(m) > max {
					max = len(m)
				}
			}
		}
		mt[i-1] = max
	}
	fmt.Println(mt)
	return mt
}

func findPresOrSufs(str string, pre bool) map[string]int {
	if len(str) == 1 {
		return map[string]int{}
	}
	rlt := map[string]int{}
	if pre {
		for i := 1; i < len(str); i++ {
			rlt[str[:i]] = 1
		}
	} else {
		for i := 1; i < len(str); i++ {
			rlt[str[i:]] = 1
		}
	}
	return rlt
}
func main() {
	//origin := "BBC ABCDAB ABCDABCDABDE"
	find := "ABCDABD"
	matchTable(find)
	//kmp(origin, find)
}
