package main

import "fmt"
import "flag"
import "time"

func DoWork() {
	time.Sleep(500 * time.Millisecond)
}

func main() {
	maxNbConcurrentGoroutines := flag.Int("maxNbConcurrentGoroutines", 5, "the number of goroutines that are allowed to run concurrently")
	nbJobs := flag.Int("nbJobs", 100, "the number of jobs that we need to do")
	flag.Parse()

	concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	for i := 0; i < *maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	done := make(chan bool)
	waitForAllJobs := make(chan bool)

	go func() {
		for i := 0; i < *nbJobs; i++ {
			<-done
			concurrentGoroutines <- struct{}{}
		}
		waitForAllJobs <- true
	}()

	for i := 0; i < *nbJobs; i++ {
		fmt.Printf("ID:%v:waiting to launch!\n", i)
		<-concurrentGoroutines
		fmt.Printf("ID:%v: it's my turn!\n", i)
		go func(id int) {
			DoWork()
			fmt.Printf("ID:%v: all done\n", id)
			done <- true
		}(i)
	}
	<-waitForAllJobs
}
