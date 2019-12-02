package main

import (
	"context"
	"fmt"
	"git.inke.cn/dating/server/base/dating.base.push/manager"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Task
type T struct {
	msg      []byte
	finished bool
}

// Worker
type W struct {
	isAvailable bool
	name        string
}

type HR struct {
	WorkercCnt int // 人数
	Mutex      sync.Mutex
	Sync       sync.WaitGroup
	TaskQ      chan T
	WorkerQ    chan W
}

func NewHR(workerCnt int) *HR {
	hr := &HR{
		WorkercCnt: workerCnt,
		Mutex:      sync.Mutex{},
		Sync:       sync.WaitGroup{},
		TaskQ:      make(chan T),
		WorkerQ:    make(chan W),
	}
	return hr
}

// 准备人手
func (hr *HR) Prepare() *HR {
	for i := 0; i < hr.WorkercCnt; i++ {
		hr.WorkerQ <- hr.hire()
	}
	return hr
}

// 雇人
func (hr *HR) hire() W {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var bit = 5
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	name := []byte{}
	for i := 0; i < bit; i++ {
		name = append(name, charset[r.Intn(len(charset))])
	}
	w := W{
		isAvailable: false,
		name:        string(name),
	}
	return w
}

// hr开始执行，让工人开始工作
func (hr *HR) Run() {
	for msg := range hr.TaskQ {
		//for i := 0; i < hr.WorkercCnt; i++ {
		//	go func() {
		for worker := range hr.WorkerQ {
			go worker.DoWork(msg)
		}
		//}()
		//}

	}
}

// hr接活
func (hr *HR) AddTask(job []byte) {
	t := &T{
		msg:      job,
		finished: false,
	}
	hr.TaskQ <- *t
}

func (w *W) DoWork(job T) {
	fmt.Println(string(job.msg))
}

func main() {
	// get kafka msg
	//taskQ := make(chan T)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//初始化一个4个工人的hr单位，并且揽活的队列为taskQ
	hr := NewHR(runtime.NumCPU())
	hr = hr.Prepare()
	hr.Run()
	//go func() {
	//	for {
	//		select {
	//		case task := <-taskQ:
	//		}
	//	}
	//}()

	for msg := range manager.GetMsgChan(ctx) {
		hr.AddTask(msg)
	}

}
