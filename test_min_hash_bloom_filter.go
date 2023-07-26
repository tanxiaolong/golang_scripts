package main

import (
    "fmt"
    "github.com/tylertreat/BoomFilters"
)

func main() {
    bag1 := []string{"bill", "alice", "frank", "bob", "sara", "tyler", "james"}
	bag2 := []string{"bill", "alice", "frank", "bob", "sara"}
	
	fmt.Println("similarity", boom.MinHash(bag1, bag2))
}
