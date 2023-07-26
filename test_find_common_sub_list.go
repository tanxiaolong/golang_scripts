// go 1.10

package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Val  int
}

func commonList(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}

}

func GenLink() (*Node, *Node) {
	head1 := &Node{}
	f1 := &Node{}
	s1 := &Node{}
	t := &Node{}
	fo := &Node{}

	head1.Val = 1
	f1.Val = 2
	s1.Val = 3
	t.Val = 4
	fo.Val = 5

	t.Next = fo
	s1.Next = t
	f1.Next = s1
	head1.Next = f1

	head2 := &Node{}
	f2 := &Node{}
	s2 := &Node{}
	t2 := &Node{}

	head2.Val = 10
	f2.Val = 20
	s2.Val = 30

	t2.Next = fo
	s2.Next = t
	f2.Next = s2
	head2.Next = f2

	return head1, head2
}

func LenLink(head *Node) uint {
	if head == nil {
		return 0
	}
	length := 1
	for head.Next != nil {
		length++
	}
	return length
}

func TraLink(head *Node) {
	for head != nil && head.Next != nil {
		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
}

func main() {
	h1, h2 := GenLink()
	fmt.Printf("h1: ")
	TraLink(h1)
	fmt.Printf("h2: ")
	TraLink(h2)
}
