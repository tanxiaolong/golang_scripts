package main

import (
	"fmt"
)
type ListNode2 struct {
	Val int
	Next *ListNode2
}
func main(){
	l1 :=&ListNode2{
		1,&ListNode2{
			2,&ListNode2{
				5,nil,
			},
		},
	}
	l2 := &ListNode2{
		1, &ListNode2{
			3, &ListNode2{
				4, nil,
			},
		},
	}
	list := mergeTwoLists(l1,l2)
	for list!=nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}


func mergeTwoLists(l1 *ListNode2,l2 *ListNode2) *ListNode2{
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	start := &ListNode2{}
	node := start
	for l1!= nil  && l2 != nil {
		start.Next = &ListNode2{}
		start = start.Next
		if l1.Val > l2.Val{
			start.Val,l2 = l2.Val,l2.Next
		}else if l1.Val < l2.Val{
			start.Val,l1 = l1.Val,l1.Next
		}else {
			start.Val = l1.Val
			start.Next = &ListNode2{Val:l2.Val}
			l1,l2 = l1.Next,l2.Next
			start = start.Next
		}

	}

	if l1 != nil {
		start.Next = l1
	}
	if l2!= nil{
		start.Next = l2
	}

	return node
}