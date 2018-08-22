package main

// backgroud:
// 给wordpress用来着
// i want to monitor a file which i append a writeStr to whether has been modified
// i monitor it alone with a modification at the same time.
import (
	"bufio"
	"fmt"
	"github.com/howeyc/fsnotify"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	previousOffset int64  = 0
	containStr     string = "user_defined_functions"
	writeStr       string = "@include(get_theme_root().'/user_defined_functions/functions.php');\n"
)

// thx to code author: "https://www.socketloop.com/tutorials/golang-simulate-tail-f-or-read-last-line-from-log-file-example"
func readLastLine(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	// we need to calculate the size of the last line for file.ReadAt(offset) to work

	// NOTE : not a very effective solution as we need to read
	// the entire file at least for 1 pass :(

	lastLineSize := 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		lastLineSize = len(line)
	}

	fileInfo, err := os.Stat(filename)

	// make a buffer size according to the lastLineSize
	buffer := make([]byte, lastLineSize)

	// +1 to compensate for the initial 0 byte of the line
	// otherwise, the initial character of the line will be missing

	// instead of reading the whole file into memory, we just read from certain offset

	offset := fileInfo.Size() - int64(lastLineSize+1)
	numRead, err := file.ReadAt(buffer, offset)

	if previousOffset != offset {

		// print out last line content
		buffer = buffer[:numRead]
		//fmt.Printf("%s \n", buffer)

		previousOffset = offset
	}
	return string(buffer)

}
func main() {
	path := os.Args[1]
	fmt.Println(path)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				// get the last line of the file, that line should be same with writeStr
				lastLine := readLastLine(path)
				// if the line do contains that key word of writeStr, then it not modified
				// or, it don't have that line, and i know that the file has been updated
				// then i append the writeStr to the end of the file.
				if !strings.Contains(lastLine, containStr) { 
					f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
					defer f.Close()
					if err != nil {
						fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
					} else {
						n, _ := f.Seek(0, os.SEEK_END)
						_, err = f.WriteAt([]byte(writeStr), n)
						fmt.Println(err)
					}

					fmt.Println(time.Now().Local())

				}
				fmt.Println(lastLine)
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Watch(path)

	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}
