package main

import "fmt"
import "unsafe"
import "sync/atomic"

type WaitGroup struct {

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}

func (wg *WaitGroup) State() (statep *uint64, semap *uint32) {
	//fmt.Println(&wg.state1)
	//fmt.Println(unsafe.Pointer(&wg.state1))
	//fmt.Println(uintptr(unsafe.Pointer(&wg.state1)))
	//fmt.Println(&wg.state1[0], &wg.state1[1], &wg.state1[2])
	//fmt.Println((*uint64)(unsafe.Pointer(&wg.state1[1])))
	fmt.Println("state1:", wg.state1)
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
		return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
	} else {
		return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
	}
}

func (wg *WaitGroup) Add(delta int) {
	statep, semap := wg.State()
	state := atomic.LoadUint64(statep)
	fmt.Println(state)
	state = atomic.AddUint64(statep, uint64(delta)<<32)
	fmt.Println(state, uint64(delta)<<32)
	v := int32(state >> 32)
	fmt.Println("vvv:", v)
	w := uint32(state)
	fmt.Println("www:", w)
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	if v > 0 || w == 0 {
		return
	}
	// This goroutine has set counter to 0 when waiters > 0.
	// Now there can't be concurrent mutations of state:
	// - Adds must not happen concurrently with Wait,
	// - Wait does not increment waiters if it sees counter == 0.
	// Still do a cheap sanity check to detect WaitGroup misuse.
	if *statep != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// Reset waiters count to 0.
	*statep = 0
	for ; w != 0; w-- {
		fmt.Println("sema:", semap)
		//runtime_Semrelease(semap, false, 0)
	}
}

func main() {
	wg := &WaitGroup{}
	state, sema := wg.State()
	fmt.Printf("1: %+v, %+v\n", state, sema)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		state, sema = wg.State()
		fmt.Printf("%d: %+v, %+v\n", i, state, sema)
	}

}
