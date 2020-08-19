
package main

import (
    "fmt"
    "github.com/tylertreat/BoomFilters"
)

func main() {
    // We could also use boom.NewUnstableBloomFilter or boom.NewPartitionedBloomFilter.
    bf := boom.NewBloomFilter(1000, 0.01)
    
    bf.Add([]byte(`a`))
    if bf.Test([]byte(`a`)) {
        fmt.Println("contains a")
    }
    
    if !bf.TestAndAdd([]byte(`b`)) {
        fmt.Println("doesn't contain b")
    }
    
    if bf.Test([]byte(`b`)) {
        fmt.Println("now it contains b!")
    }
    
    // Restore to initial state.
    bf.Reset()
}
