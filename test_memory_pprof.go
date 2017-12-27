package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sync"
	"fmt"
)

func bigBytes() *[]byte {
	s := make([]byte, 10000000000)
	return &s
}

func main() {

	var wg sync.WaitGroup
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	pprof.WriteHeapProfile(os.Stdout)
	fmt.Println()
	fmt.Printf("goronties:%+v\n",pprof.Lookup("goroutine").WriteTo(os.Stdout,1))
	//fmt.Printf("goronties:%+v\n",pprof.Lookup("threadcreate").WriteTo(os.Stdout,2))
	//fmt.Printf("goronties:%+v\n",pprof.Lookup("heap").WriteTo(os.Stdout,2))
	//fmt.Printf("goronties:%+v\n",pprof.Lookup("block").WriteTo(os.Stdout,2))
	wg.Add(1)
	wg.Wait()
}
