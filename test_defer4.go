package main

import "fmt"

func test()(i int){
  defer func(){
    i++
  }()
  return 1
}

func test2() int {
  i := 1 
  defer func(){
    i++
  }()
  return i
}
func main(){
  fmt.Println(test(),test2())
}
