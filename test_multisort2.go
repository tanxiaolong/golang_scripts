package main

import (
	"fmt"
	"sort"
)

type ItemsT []ItemT

type ItemT struct {
	Order int
	Age   int
	Name  string
}

type Bucket struct {
	Slice []interface{}               //承载以任意结构体为元素构成的Slice
	By    func(a, b interface{}) bool //排序规则函数,当需要对新的结构体slice进行排序时，只需定义这个函数即可
}

/*
定义三个必须方法的准则：接收者不能为指针
*/
func (t Bucket) Len() int { return len(t.Slice) }

func (t Bucket) Swap(i, j int) { t.Slice[i], t.Slice[j] = t.Slice[j], t.Slice[i] }

func (t Bucket) Less(i, j int) bool { return t.By(t.Slice[i], t.Slice[j]) }

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
	c := &Bucket{}
	for i := range tmp {
		c.Slice = append(c.Slice, &tmp[i])
	}
	c.By = func(a, b interface{}) bool {
		if a.(*ItemT).Age < a.(*ItemT).Age {
			return true
		}
		if a.(*ItemT).Age > a.(*ItemT).Age {
			return false
		}
		return a.(*ItemT).Order < a.(*ItemT).Order
	}
	fmt.Printf("%+v\n", *c.Slice)
	sort.Sort(c)
	fmt.Printf("%+v\n", c.Slice)
	fmt.Println(c.Slice)
}
