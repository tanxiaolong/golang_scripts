package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

const DIR = "/script/golang/"

func main() {
	dir_list, err := ioutil.ReadDir(os.Getenv("HOME") + DIR)
	if err != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dir_list {
		fileName := v.Name()
		if path.Ext(fileName) == ".go" {
			cmd := exec.Command("go", "fmt", fileName)
			if err := cmd.Run(); err == nil {
				fmt.Println(fileName, " fmt Done.")
			} else {
				fmt.Println(err)
				fmt.Println(fileName, " fmt failure.")
			}
		}
	}
}
