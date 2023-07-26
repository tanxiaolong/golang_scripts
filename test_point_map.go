package main
import "fmt"
type Node struct {
    Val int
    Next *Node
    Random *Node
}

func main(){
	m := map[*Node]*Node{}

	t1 := &Node{
		Val: 1,
		Next: &Node{},
		Random: &Node{},
	}
	m[t1] = t1
	fmt.Println(m[t1])
}
