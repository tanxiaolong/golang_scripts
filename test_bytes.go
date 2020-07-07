package main

import "fmt"

func main() {
	s := "a good   example"
	bytes := []byte(s)
	p := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != ' ' {
			wordsStart := i
			wordsStop := 0
			for j := i + 1; j < len(bytes); j++ {
				p = j
				if bytes[j] == ' ' {
					wordsStop = j
					i = j //找到stop后，更新下一个start
					break
				}
			}
			fmt.Println("p:", p)
			if p == len(bytes)-1 {
				wordsStop = len(bytes)
			}
			revers(bytes[wordsStart:wordsStop])
			if wordsStop == len(bytes) {
				break
			}
		} else {
			// 这里有点问题
			// 1、只考虑到了单词之间两个空格的情况，
			// 2、没法变短bytes
			fmt.Println(i, p)
			if bytes[i] == bytes[p] {
				for m := i; m < len(bytes)-1; m++ {
					bytes[m] = bytes[m+1]
				}
			}
		}
	}
	rlt := string(revers(bytes))
	i, j := 0, len(rlt)-1
	for {
		if rlt[i] == ' ' {
			i++
		}
		if rlt[j] == ' ' {
			j--
		}
		if rlt[i] != ' ' && rlt[j] != ' ' {
			break
		}
	}
	rlt = rlt[i : j+1]
	fmt.Println(rlt)
}

func revers(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
