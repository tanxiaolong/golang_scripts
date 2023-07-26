package main

import "fmt"
import "os"
import "log"
import "time"

func main() {
	i := 0
	for {
		go writeToFile(fmt.Sprintf("%d\n", i))
		i++
	}
	writeToFile("start\n")
	time.Sleep(5)
}

func writeToFile(msg string) {
	f, err := os.OpenFile("./ttttttttttttttttt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}
