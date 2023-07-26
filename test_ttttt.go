package main
import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)
func main(){
	var x int
	defer func(){
		fmt.Println("defer: ",x)
	}()
	threads := runtime.GOMAXPROCS(0)
	fmt.Println("threads cnt: ", threads)
	for i:=0;i<threads;i++{
		go func(){
			x++
		}()
	}
	//runtime.GC()
	time.Sleep(time.Second)
	fmt.Println(x)
	
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs: %d\n", numCPU)

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	for i := 0; i < numCPU; i++ {
		processor := &stats.PauseNs[i]
		queueLen := *(*uint32)(unsafe.Pointer(processor))
		fmt.Printf("Local queue length of Processor %d: %d\n", i, queueLen)
	}
}
