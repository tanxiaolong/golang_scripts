package main

import (
	_"fmt"
	"github.com/rjeczalik/notify"
	"log"
	"os"
	"os/exec"
	"strings"
	"fmt"
)

func main() {
	route := os.Args[1]
	fmt.Println(route)
	
	//os.Exit(0)
	c := make(chan notify.EventInfo, 1)
	//if err := notify.Watch("/home/tanxiaolong/chain-admin...", c, notify.All); err != nil {
	if err := notify.Watch("./...", c, notify.All); err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-c:
				path := ev.Path()
				originPrefix := "/home/tanxiaolong/chain-admin"
				prefix := "/letv/www/chain-admin-txl"
				newPath := strings.Replace(path, originPrefix, prefix, -1)
				_, err := os.Stat(newPath)
				if err == nil { // if file exists
					_, err := exec.Command("rm", "-rf", newPath).Output()
					if err != nil {
						log.Println("err:", err)
					}
					_, err = exec.Command("cp", path, newPath).Output()
					if err != nil {
						log.Println("err:", err)
					}
					_, err = exec.Command("chown", "-R", "www:www", newPath).Output()
					if err != nil {
						log.Println("err:", err)
					}
					log.Println("event:", ev.Path())
					log.Println("asd:", ev.Event())
					log.Println("替换成功")
				}else{
					log.Println(newPath,"路径不存在")
				}

			}

		}
	}()
	<-done
}
