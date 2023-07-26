package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 14042847
	b := 16846487
	fmt.Printf("message_info_%d", GenIndex(int64(a+b), 32))
	fmt.Println()
	fmt.Printf("info_%d", GenIndex(int64(GenPartition(int64(a+b), 32)), 128))
	fmt.Println()
	fmt.Printf("union_id: %s", GetUnionId(a, b))
}

func GenIndex(seed int64, maxIndex uint32) uint32 {
	r := rand.New(rand.NewSource(seed))
	idx := r.Uint32() % maxIndex
	return idx
}

func GenPartition(seed int64, maxIndex uint32) uint32 {
	r := rand.New(rand.NewSource(seed))
	idx := r.Uint32() / maxIndex
	return idx
}

func GetUnionId(ownerId int, peerId int) string {
	var unionId string
	if ownerId < peerId {
		unionId = fmt.Sprintf("%d:%d", ownerId, peerId)
	} else {
		unionId = fmt.Sprintf("%d:%d", peerId, ownerId)
	}
	return unionId
}
