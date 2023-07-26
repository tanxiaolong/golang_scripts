package main

import "fmt"

//import "sort"
import "sync"
import "time"

func main() {
	wg := sync.WaitGroup{}
	a := []int{}
	mutex := sync.Mutex{}
	start := time.Now()
	for i := 0; i < 3000000; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
			a = append(a, i)
		}(i)
	}
	wg.Wait()

	gap := time.Since(start)
	fmt.Println(gap)

	start = time.Now()
	for i := 0; i < 3000000; i++ {
		a = append(a, i)
	}
	gap = time.Since(start)
	fmt.Println(gap)

	c := make(chan struct{})
	s := NewService(100, func() { c <- struct{}{} })

	// 2. 起n个goroutine不断执行增加操作
	start = time.Now()
	n := 3000000
	var wg2 sync.WaitGroup
	wg2.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			s.Add(a)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	s.Close()

	<-c

	gap = time.Since(start)
	fmt.Println(gap)

	// 3. 校验所有结果是否都被添加上
	//fmt.Println("done len:", len(s.Slice()))
}

// active object对象
type Service struct {
	channel chan int `desc:"即将加入到数据slice的数据"`
	data    []int    `desc:"数据slice"`
}

// 新建一个size大小缓存的active object对象
func NewService(size int, done func()) *Service {
	s := &Service{
		channel: make(chan int, size),
		data:    make([]int, 0),
	}

	go func() {
		s.schedule()
		done()
	}()
	return s
}

// 把管道中的数据append到slice中
func (s *Service) schedule() {
	for v := range s.channel {
		s.data = append(s.data, v)
	}
}

// 增加一个值
func (s *Service) Add(v int) {
	s.channel <- v
}

// 管道使用完关闭
func (s *Service) Close() {
	close(s.channel)
}

// 返回slice
func (s *Service) Slice() []int {
	return s.data
}
