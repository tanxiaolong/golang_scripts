package main

import (
	_ "bufio"
	_ "io/ioutil"
	"os"
)

func main() {
	content := []byte("asd")
	write(content)
}

func write(content []byte) (err error) {
	f, err := os.OpenFile("filePath", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}
	// 性能较差，但是写有报障
	return f.Sync()
}