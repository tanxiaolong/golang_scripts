package main

import (
	"fmt"
	"os"
	"golang.org/x/sys/unix"
)

func main() {

	var cpuset unix.CPUSet
	err := unix.SchedGetaffinity(0, &cpuset)
	if err != nil {
		fmt.Printf("numcpus: cannot get affinity: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", cpuset.Count())
}