package main

import "fmt"
import "time"

func main(){
	m := map[int]int{1:1}
	for i:=0;i<1000;i++{
		go func(){
			fmt.Println(m[1])
		}()
	}
	time.Sleep(5*time.Second)
}
