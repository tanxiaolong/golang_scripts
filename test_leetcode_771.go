package main

import "log"

func main(){
	J := "aA"
	S := "aAAbbbb"
	n := numJewelsInStones(J,S)
	log.Print(n)
}


func numJewelsInStones(J string, S string) int {
	if J == "" || S == "" {
		return 0
	}
	tmp := map[string]int{}
	for _,c := range J {
		tmp[string(c)] = 1
	}
	rlt := 0
	for _,c := range S {
		if _,exists := tmp[string(c)];exists{
			rlt ++
		}
	}
	return rlt
}
