package main

import (
	"fmt"
)

func main() {

	illegalWord1 := "干脆面"
	illegalWord2 := "方便面"
	illegalWord3 := "干死你"

	treeMap := map[rune]DFANode{}
	t := initDFA(treeMap, illegalWord1)
	t = initDFA(t, illegalWord2)
	t = initDFA(t, illegalWord3)

	test := "我非常爱吃干脆面和方便面"
	fmt.Println(isLegal(t, test))
}

type DFANode struct {
	next    map[rune]DFANode
	notlast bool
}

func isLegal(treemap map[rune]DFANode, str string) (bool, string) {

	return true, ""
}

func initDFA(treemap map[rune]DFANode, sensitiveWord string) map[rune]DFANode {
	rlt := map[rune]DFANode{}
	if len(sensitiveWord) == 0 {
		return rlt
	}
	chars := []rune(sensitiveWord)
	length := len(chars)
	for i := range chars {
		letter := chars[i]
		if _, exists := treemap[letter]; !exists {
			treemap[letter] = DFANode{}
		}

	}
	return rlt
}
