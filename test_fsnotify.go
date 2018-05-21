package main

import (
	"log"
	_"fmt"
	"github.com/howeyc/fsnotify"
)

func main(){
        watcher, err:= fsnotify.NewWatcher()
        if err!=nil {
                log.Fatal(err)
        }
        done := make(chan bool)

        go func(){
                for {
                        select {
                                case ev := <- watcher.Event:
                                        log.Println("event:",ev)
                                case err := <- watcher.Error:
                                        log.Println("error:",err)
                        }
                }
        }()

        err = watcher.Watch("/home/tanxiaolong/scripts")
        if err != nil{
                log.Fatal(err)
        }

        <-done

        watcher.Close()
}

