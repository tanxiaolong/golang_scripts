//go 1.10
package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   interface{}
}

/*
    root
  l1    r1
l2  r2
   l3

output:
a,b,c,d,e,nil,nil,nil,nil,f,nil

*/
func GenTree() *Node {
	root := &Node{} // a
	l1 := &Node{}   // b
	r1 := &Node{}   // c
	l2 := &Node{}   // d
	r2 := &Node{}   // e
	l3 := &Node{}   // f
	root.Left = l1
	root.Right = r1
	root.Val = 0
	l1.Left = l2
	l1.Right = r2
	l1.Val = 1
	r1.Left = nil
	r1.Right = nil
	r1.Val = 2
	l2.Left = nil
	l2.Right = nil
	l2.Val = 3
	r2.Left = l3
	r2.Right = nil
	r2.Val = 4
	l3.Left = nil
	l3.Right = nil
	l3.Val = 5

	return root
}

func LayerTraverse(root *Node) {
	q := list.New()
	if root == nil {
		return
	}
	q.PushBack(root)
	flag := &Node{}
	flag.Left = nil
	flag.Right = nil
	flag.Val = "#"
	q.PushBack(flag)
	for q.Len() > 0 {
		ele := q.Front()
		v := ele.Value.(*Node)
		fmt.Println(v.Val)
		switch v.Val.(type) {
		case string:
			fmt.Println()
			q.Remove(ele)
		case int:
			fmt.Printf("%d ", v.Val.(int))
			q.Remove(ele)
			continue
		}
		if v.Left != nil {
			q.PushBack(v.Left)
		}
		if v.Right != nil {
			q.PushBack(v.Right)
		}
		q.PushBack("#")
	}
}

func main() {
	root := GenTree()
	LayerTraverse(root)
}
