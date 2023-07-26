package main

import "fmt"

import "time"
import "runtime"

func main() {
	//print3()
	print2()
}

func print() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 101; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
	}()
	go func() {
		for i := 0; i < 101; i++ {
			if i%2 != 0 {
				ch <- i
			}
		}
	}()
	for {
		select {
		case d := <-ch:
			fmt.Println(d)
		}
	}
}

func print2() {
	c := make(chan int)
	go func() {
		for i := 1; i < 101; i++ {
			c <- 1
			//奇数
			if i%2 == 1 {
				fmt.Println("线程a打印:", i)
			}
		}
	}()
	go func() {
		for i := 1; i < 101; i++ {
			<-c
			//偶数
			if i%2 == 0 {
				fmt.Println("线程b打印:", i)
			}
		}
	}()
	time.Sleep(3 * time.Second)
}

func print3() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 1; i < 101; i++ {
			//奇数
			if i%2 == 1 {
				fmt.Println("线程1打印:", i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 1; i < 101; i++ {
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:", i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	time.Sleep(3 * time.Second)
}
