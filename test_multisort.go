package main

import (
	"fmt"
	"sort"
)

type ItemT struct {
	Order int
	Age   int
	Name  string
}

type ItemsT []ItemT

func (p ItemsT) Len() int {
	return len(p)
}

//func (p ItemsT) Less(i, j int) bool {
//	return p[i].Order < p[j].Order
//}
/*
func (p ItemsT) Less(i, j int) bool {
	if p[i].Order < p[j].Order {
		return true
	}
	if p[i].Order > p[j].Order {
		return false
	}

	return p[i].Age < p[j].Age
}
*/
func (p ItemsT) Less(i, j int) bool {
	if p[i].Age < p[j].Age {
		return true
	}
	if p[i].Age > p[j].Age {
		return false
	}

	return p[i].Order < p[j].Order
}
func (p ItemsT) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *ItemsT) Sort() {
	sort.Sort(p)
}

// 主函数
func main() {
	tmp := ItemsT{
		{
			Order: 13,
			Name:  "F",
			Age:   43,
		},
		{
			Order: 10,
			Name:  "A",
			Age:   28,
		},
		{
			Order: 5,
			Name:  "B",
			Age:   57,
		},
		{
			Order: 5,
			Name:  "C",
			Age:   43,
		},
		{
			Order: 20,
			Name:  "D",
			Age:   16,
		},
		{
			Order: 1,
			Name:  "E",
			Age:   31,
		},
	}
	tmp.Sort()
	fmt.Printf("%+v", tmp)
}
