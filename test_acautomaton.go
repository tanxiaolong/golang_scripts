package main

import (
	"fmt"
	"github.com/zheng-ji/goAcAutoMachine"
)

func main() {
	ac := goAcAutoMachine.NewAcAutoMachine()
	ac.AddPattern("红领巾")
	ac.AddPattern("祖国")
	ac.AddPattern("花朵")
	ac.Build()

	content := "我是红领巾，祖国的未来"
	results := ac.Query(content)
	for _, result := range results {
		fmt.Println(result)
	}
}
