package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := map[string]interface{}{
		"a": "b",
	}

	b := a
	a["b"] = "a"
	fmt.Println(a, b)

	c := map[string]interface{}{}
	for k, v := range a {
		c[k] = v
	}
	a["c"] = "c"

	fmt.Println(a, c)
	fmt.Println("---------------------------")
	//persons := make(map[string]int)
	persons := map[string]int{}
	persons["张三"] = 19

	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modify(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func modify(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}
