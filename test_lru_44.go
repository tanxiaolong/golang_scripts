package main

import "lrucache"
import "fmt"

func main() {
	cache := lrucache.NewLruCache(10)
	fmt.Println(cache.Set("a", "c"))
	fmt.Println(cache.Get("a"))
	fmt.Println(cache.GetCap())
	fmt.Println(cache.GetUsed())
	fmt.Println(cache.Traversal())
	fmt.Println(cache)
	fmt.Println(cache.ResetLruCache())
}
