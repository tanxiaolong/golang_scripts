package main

//import "fmt"

func main(){
	// string 线程不安全
	//str := "1"
	var a bool
	go func(){
		a = true
		//fmt.Println(str)
	}()

	go func(){
		a = false
	}()


}
