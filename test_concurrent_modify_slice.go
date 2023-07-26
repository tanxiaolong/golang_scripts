package main

import (
	"fmt"
	"time"
)
import "sync"

func ConcurrentModify(rlt []string) []string {

	target := make([]string, len(rlt))
	wg := sync.WaitGroup{}
	for i, item := range rlt {
		item := item
		wg.Add(1)
		go func(index int, origin *[]string) {
			defer wg.Done()
			target[index] = item + "abc"
		}(i, &rlt)
	}
	wg.Wait()
	return target
}

func SeqModify(rlt []string) []string {
	target := make([]string, len(rlt))
	for i, item := range rlt {
		item := item
		target[i] = item + "abc"
	}
	return target
}

func main() {

	origin := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println("origin:", origin)
	startAt := time.Now()
	target := ConcurrentModify(origin)
	duration := time.Since(startAt)
	fmt.Println("concurrent:", duration, "target:", target)
	startAt = time.Now()
	target = SeqModify(origin)
	duration = time.Since(startAt)
	fmt.Println("seq:", duration, "target:", target)
	fmt.Println("aaaa: %v", duration.Seconds())
}
