package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	myPath := "/home/tanxiaolong/script/golang/test"
	fmt.Println(listPath(myPath))
}

func listPath(path string) []string {
	filename := []string{}
	m := map[string]int{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		filename = append(filename, info.Name())
		fmt.Println(info.Mode())
		m[info.Name()] = 1
		return nil
		if _, ok := m[info.Name()]; !ok && info.IsDir() {
			listPath(path)
		}
		return nil
	})
	return filename
}
