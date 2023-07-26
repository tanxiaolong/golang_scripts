package main

import "sync/atomic"

func main(){
	s := &pad{
		x: 0,
		_: [56]byte{},
		y: 0,
		z: 0,
	}
	s.increase()
}

// pad 结构的 x y z 会被并发的执行原子操作
type pad struct {
	x uint64 // 8byte
	_ [56]byte
	y uint64 // 8byte
	_ [56]byte
	z uint64 // 8byte
	_ [56]byte
}

func (s *pad) increase() {
	atomic.AddUint64(&s.x, 1)
	atomic.AddUint64(&s.y, 1)
	atomic.AddUint64(&s.z, 1)
}