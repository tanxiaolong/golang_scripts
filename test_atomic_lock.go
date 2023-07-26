package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	num := int32(10000000)
	testAtomic(num)
	testLock(num)
}
func testLock(count int32) {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	var a int32
	t1 := time.Now()
	for i := int32(0); i < count; i++ {
		wg.Add(1)
		go func(i int32) {
			defer func() {
				wg.Done()
				lock.Unlock()
			}()
			lock.Lock()
			a = i + count
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(t1))
	fmt.Println("lock", a)
}
func testAtomic(count int32) {
	var a int32
	wg := sync.WaitGroup{}
	t1 := time.Now()
	for i := int32(0); i < count; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			atomic.SwapInt32(&a, i+count)
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(t1))
	fmt.Println("atomic", a)
}
