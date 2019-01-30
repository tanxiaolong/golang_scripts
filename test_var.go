package main
import (
	"fmt"
)

func main(){
 	// 这里只是分配了一个指针地址
	var n *int // 申请了指针地址
	// 如果没有下面这句——给地址分配内存空间，那么数据放哪儿呢？没地儿放就会报错。
	n = new(int) // 赋予了内存空间
	*n = 10
	fmt.Println(*n)


	// e.g.2
	var b = new(int) 
	//b := new(int) // := 操作直接给了地址和内存空间
        fmt.Println(*b)

	// e.g.3
        fmt.Println(*c)
}
