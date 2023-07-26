package main

import "fmt"
import "sort"

func main() {
	a := []string{"3", "2", "1", "5", "4", "8", "0"}
	sort.Slice(a, func(i, j int) bool {
		if a[i] <= "3" {
			a[i] = a[i]+"km"
			fmt.Println(a)
		}
		return a[i] < a[j]
	})
	fmt.Println(a)
}
