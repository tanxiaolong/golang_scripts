package main

import (
	"fmt"
	"time"
)

func main() {
	pool := NewWorker(100)
	for i:=0;i<50;i++ {
		fmt.Println("add job ",i)
		pool.AddTask(func() {
			fmt.Println("done",i)
		})
	}
	time.Sleep(10)
}

// NewWorker 创建一个协程池
func NewWorker(maxWorkers int) *WorkerPool {
	if maxWorkers < 1 {
		maxWorkers = 1
	}

	pool := &WorkerPool{
		workerCont:  0,
		maxWorkers:  maxWorkers,
		taskQueue:   make(chan func(), 1),
		ready:       make(chan int, 1),
		workerQueue: make(chan func(), 1),
	}
	go pool.dispatch()
	go pool.runWorker()
	return pool
}

// WorkerPool 协程池
type WorkerPool struct {
	workerCont  int
	maxWorkers  int
	taskQueue   chan func()
	ready       chan int
	workerQueue chan func()
}

// MaxSize 协程池最大大小
func (wp *WorkerPool) MaxSize() int {
	return wp.maxWorkers
}

// AddTask 增加任务
func (wp *WorkerPool) AddTask(task func()) {
	if task != nil {
		fmt.Println("added")
		wp.taskQueue <- task
	}
}

// 主要用于消费等，这里会阻塞主线程
func (wp *WorkerPool) dispatch() {
	end := false
	for {
		if end {
			break
		}
		select {
		case task, ok := <-wp.taskQueue:
			fmt.Println("out",ok)
			if !ok {
				end = true
				break
			}
			//task()
			// 看看是否有闲置队列，如果有，就使用闲置队列
			select {
			case <-wp.ready:
				fmt.Println("send job")
				wp.workerQueue <- task
			default:
				if wp.workerCont < wp.maxWorkers {
					fmt.Println(wp.workerCont,wp.maxWorkers)
					wp.workerCont++
					go wp.runWorker()
				//} else {
				//	fmt.Println("task in q")
				//	wp.workerQueue <- task
				//	fmt.Println("task in q count: ",len(wp.workerQueue))
				}
				wp.workerQueue <- task

			}
		//default:
		//	fmt.Println("default outer")
		}
	}
}

func (wp *WorkerPool) runWorker() {
	fmt.Println("job count: ",len(wp.workerQueue))
	for task := range wp.workerQueue {
		task()
		wp.ready <- 1
	}
}
