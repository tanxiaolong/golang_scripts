package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

func main() {
	fmt.Println("vim-go")

	//l14 := &ListNode{
	//	Val:  9,
	//	Next: nil,
	//}
	//l13 := &ListNode{
	//	Val:  8,
	//	Next: l14,
	//}
	l12 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l11 := &ListNode{
		Val:  9,
		Next: l12,
	}

	l1 := &ListNode{
		Val:  8,
		Next: l11,
	}

	//l22 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//l21 := &ListNode{
	//	Val:  9,
	//	Next: nil,
	//}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	tmp := addTwoNumbers(l1, l2)
	for tmp != nil {
		fmt.Println("???:", tmp.Val)
		tmp = tmp.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	llen1 := length(l1)
	llen2 := length(l2)
	fmt.Println(llen1, llen2)
	var rlt *ListNode
	if llen1 > llen2 {
		rlt = add(l1, l2)
	} else if llen1 < llen2 {
		rlt = add(l2, l1)
	} else {
		rlt = add(l1, l2)
	}

	return rlt
}
func add(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp int
	tListNode := l1
	var lastL1 *ListNode
	for l1 != nil {
		if l2 == nil {
			break
		}
		fmt.Println(l1.Val, l2.Val, tmp)
		tmp = l1.Val + l2.Val + tmp
		if tmp >= 10 {
			tmp = tmp % 10
			l1.Val = tmp
			lastL1 = l1
			l1 = l1.Next
			l2 = l2.Next
			tmp = 1
			continue
		}
		l1.Val = tmp
		lastL1 = l1
		l1 = l1.Next
		l2 = l2.Next
		tmp = 0
	}
	fmt.Println("123:", tmp)
	for tmp == 1 {
		if l1 == nil {
			fmt.Println(4444)
			l1 = lastL1

			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			tmp = 0
		} else {
			fmt.Println(3333333)
			l1.Val = l1.Val + tmp
			if l1.Val >= 10 {
				tmp = l1.Val % 10
				l1.Val = tmp
				if l1.Next == nil {
					l1.Next = &ListNode{
						Val:  1,
						Next: nil,
					}
					break
				}
				tmp = 1
				l1 = l1.Next
			}
			tmp = 0
		}
	}
	return tListNode
}

func length(l *ListNode) int {
	var i int
	for l.Next != nil {
		i++
		l = l.Next
	}
	return i
}
