//go 1.10
package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  int
}

func GenLink() *Node {
	head := &Node{}
	f := &Node{}
	s := &Node{}
	t := &Node{}

	head.Val = 1
	f.Val = 2
	s.Val = 3
	t.Val = 4

	s.Next = t
	f.Next = s
	head.Next = f

	return head
}

func TraLink(head *Node) {
	for head != nil && head.Next != nil {
		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}

func RevrLink(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head
	head = head.Next
	tmp.Next = nil
	for head.Next != nil {
		n := head.Next
		head.Next = tmp
		tmp = head
		head = n
	}

	head.Next = tmp
	return head
}
func main() {
	head := GenLink()
	TraLink(head)
	rLink := RevrLink(head)
	TraLink(rLink)
}
