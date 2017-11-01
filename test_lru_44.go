package main

import "lrucache"
import "fmt"

func main(){
	cache := lrucache.NewLruCache()
	fmt.Println(cache.Set("a","c"))
	fmt.Println(cache.Get("a"))
	fmt.Println(cache.GetCap())
	fmt.Println(cache.GetUsed())
	fmt.Println(cache.Traversal())
}
