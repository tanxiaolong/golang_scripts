package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

// 全局变量（简单处理）
var p Person

func update(name string, age int) {
	// 更新第一个字段

	p = Person{
		name: name,
		age:  age,
	}
	//p.name = name
	// 加点随机性
	time.Sleep(time.Millisecond * 200)
	// 更新第二个字段
	//p.age = age
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 10 个协程并发更新
	for i := 0; i < 10; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			update(name, age)
		}()
	}
	wg.Wait()
	// 结果是啥？你能猜到吗？
	fmt.Printf("p.name=%s\np.age=%v\n", p.name, p.age)
}

type Test struct {
	//X int
	Y string
}

func test() {
	var g Test

	for i := 0; i < 1000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = Test{
				//X: 1,
				Y: "2",
			}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = Test{
				//X: 3,
				Y: "4",
			}
		}()

		wg.Wait()
		// 赋值异常判断
		//if !((g.X == 1 && g.Y == "2") || (g.X == 3 && g.Y == "4")) {
		//if !((g.X == 1) || (g.X == 3)) {
		if !((g.Y == "2") || (g.Y == "4")) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v", i, g)
			break
		}
	}

	fmt.Printf("gggg: %v\n", g)
}
