package main

import "fmt"
import "sync"

func main() {
	//all := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	all := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	uidCh := make(chan uint64, len(all))
	for _, uid := range all {
		uidCh <- uid
	}
	close(uidCh)
	sg := sync.WaitGroup{}
	// 发起多个协程并发去更新会话记录
	for i := 0; i < 5; i++ {
		sg.Add(1)
		go func(i int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("asd")
				}
				sg.Done()
			}()
			for uid := range uidCh {
				fmt.Println("work", i, ":", uid)
			}
		}(i)
	}
	sg.Wait()
}
