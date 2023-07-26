package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	b := Float32ToByte(1.23)
	fmt.Println(b)
	b = Float32ToByte(0.1)
	fmt.Println(b)
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}
